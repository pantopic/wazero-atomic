package wazero_atomic

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// Name is the name of this host module.
const Name = "pantopic/wazero-atomic"

var (
	ctxKeyMeta   = Name + `/meta`
	ctxKeyUint64 = Name + `/uint64`
)

type meta struct {
	ptrSet    uint32
	ptrID     uint32
	ptrUint64 uint32
}

type hostModule struct {
	sync.RWMutex

	module api.Module
}

type Option func(*hostModule)

func New(opts ...Option) *hostModule {
	p := &hostModule{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (h *hostModule) Name() string {
	return Name
}

func (h *hostModule) ContextCopy(dst, src context.Context) context.Context {
	dst = context.WithValue(dst, ctxKeyMeta, get[*meta](src, ctxKeyMeta))
	if v := src.Value(ctxKeyUint64); v != nil {
		dst = context.WithValue(dst, ctxKeyUint64, v.(map[uint32]map[uint64]*atomic.Uint64))
	} else {
		dst = context.WithValue(dst, ctxKeyUint64, make(map[uint32]map[uint64]*atomic.Uint64))
	}
	return dst
}

func (h *hostModule) Stop() {}

// Register instantiates the host module, making it available to all module instances in this runtime
func (h *hostModule) Register(ctx context.Context, r wazero.Runtime) (err error) {
	builder := r.NewHostModuleBuilder(Name)
	register := func(name string, fn func(ctx context.Context, m api.Module, stack []uint64)) {
		builder = builder.NewFunctionBuilder().WithGoModuleFunction(api.GoModuleFunc(fn), nil, nil).Export(name)
	}
	for name, fn := range map[string]any{
		"__atomic_uint64_add": func(u64 *atomic.Uint64, delta uint64) (new uint64) {
			return u64.Add(delta)
		},
		"__atomic_uint64_load": func(u64 *atomic.Uint64) uint64 {
			return u64.Load()
		},
		"__atomic_uint64_store": func(u64 *atomic.Uint64, val uint64) {
			u64.Store(val)
		},
		"__atomic_uint64_del": func(m map[uint64]*atomic.Uint64, id uint64) {
			h.Lock()
			delete(m, id)
			h.Unlock()
		},
	} {
		switch fn := fn.(type) {
		case func(u64 *atomic.Uint64, delta uint64) (new uint64):
			register(name, func(ctx context.Context, mod api.Module, stack []uint64) {
				meta := get[*meta](ctx, ctxKeyMeta)
				delta := readUint64(mod, meta.ptrUint64)
				u64 := h.getCtxU64(ctx, mod, meta)
				new := fn(u64, delta)
				writeUint64(mod, meta.ptrUint64, new)
			})
		case func(u64 *atomic.Uint64) uint64:
			register(name, func(ctx context.Context, mod api.Module, stack []uint64) {
				meta := get[*meta](ctx, ctxKeyMeta)
				u64 := h.getCtxU64(ctx, mod, meta)
				val := fn(u64)
				writeUint64(mod, meta.ptrUint64, val)
			})
		case func(u64 *atomic.Uint64, val uint64):
			register(name, func(ctx context.Context, mod api.Module, stack []uint64) {
				meta := get[*meta](ctx, ctxKeyMeta)
				val := readUint64(mod, meta.ptrUint64)
				u64 := h.getCtxU64(ctx, mod, meta)
				fn(u64, val)
			})
		case func(m map[uint64]*atomic.Uint64, id uint64):
			register(name, func(ctx context.Context, mod api.Module, stack []uint64) {
				meta := get[*meta](ctx, ctxKeyMeta)
				val := readUint64(mod, meta.ptrID)
				m := h.getCtxU64Set(ctx, mod, meta)
				fn(m, val)
			})
		default:
			log.Panicf("Method signature implementation missing: %#v", fn)
		}
	}
	h.module, err = builder.Instantiate(ctx)
	return
}

// InitContext retrieves the meta page from the wasm module
func (h *hostModule) InitContext(ctx context.Context, m api.Module) (context.Context, error) {
	stack, err := m.ExportedFunction(`__atomic`).Call(ctx)
	if err != nil {
		return ctx, err
	}
	meta := &meta{}
	ptr := uint32(stack[0])
	for i, v := range []*uint32{
		&meta.ptrSet,
		&meta.ptrID,
		&meta.ptrUint64,
	} {
		*v = readUint32(m, ptr+uint32(4*i))
	}
	return context.WithValue(ctx, ctxKeyMeta, meta), nil
}

func (h *hostModule) getCtxU64Set(ctx context.Context, mod api.Module, meta *meta) map[uint64]*atomic.Uint64 {
	set := readUint32(mod, meta.ptrSet)
	m := get[map[uint32]map[uint64]*atomic.Uint64](ctx, ctxKeyUint64)
	h.RLock()
	s, ok := m[set]
	h.RUnlock()
	if !ok {
		h.Lock()
		if s, ok = m[set]; !ok {
			s = make(map[uint64]*atomic.Uint64)
			m[set] = s
		}
		h.Unlock()
	}
	return s
}

func (h *hostModule) getCtxU64(ctx context.Context, mod api.Module, meta *meta) *atomic.Uint64 {
	s := h.getCtxU64Set(ctx, mod, meta)
	id := readUint64(mod, meta.ptrID)
	h.RLock()
	u64, ok := s[id]
	h.RUnlock()
	if !ok {
		h.Lock()
		if _, ok = s[id]; !ok {
			u64 = &atomic.Uint64{}
			s[id] = u64
		}
		h.Unlock()
	}
	return u64
}

func get[T any](ctx context.Context, key string) T {
	v := ctx.Value(key)
	if v == nil {
		log.Panicf("Context item missing %s", key)
	}
	return v.(T)
}

func id(m api.Module, meta *meta) uint32 {
	return readUint32(m, meta.ptrID)
}

func getSet(m api.Module, meta *meta) uint32 {
	return readUint32(m, meta.ptrSet)
}

func readUint32(m api.Module, ptr uint32) (val uint32) {
	val, ok := m.Memory().ReadUint32Le(ptr)
	if !ok {
		log.Panicf("Memory.Read(%d) out of range", ptr)
	}
	return
}

func read(m api.Module, ptrData, ptrLen, ptrMax uint32) (buf []byte) {
	buf, ok := m.Memory().Read(ptrData, readUint32(m, ptrMax))
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", ptrData, ptrLen)
	}
	return buf[:readUint32(m, ptrLen)]
}

func readUint64(m api.Module, ptr uint32) (val uint64) {
	val, ok := m.Memory().ReadUint64Le(ptr)
	if !ok {
		log.Panicf("Memory.Read(%d) out of range", ptr)
	}
	return
}

func writeUint64(m api.Module, ptr uint32, val uint64) {
	if ok := m.Memory().WriteUint64Le(ptr, val); !ok {
		log.Panicf("Memory.Read(%d) out of range", ptr)
	}
}

func writeUint32(m api.Module, ptr uint32, val uint32) {
	if ok := m.Memory().WriteUint32Le(ptr, val); !ok {
		log.Panicf("Memory.Read(%d) out of range", ptr)
	}
}
