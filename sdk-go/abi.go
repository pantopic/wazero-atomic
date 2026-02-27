package atomic

import (
	"unsafe"
)

var (
	id   uint32
	u64  uint64
	meta = make([]uint32, 2)
)

//export __atomic
func __atomic() (res uint32) {
	meta[0] = uint32(uintptr(unsafe.Pointer(&id)))
	meta[1] = uint32(uintptr(unsafe.Pointer(&u64)))
	return uint32(uintptr(unsafe.Pointer(&meta[0])))
}

//go:wasm-module pantopic/wazero-atomic
//export __atomic_uint64_add
func uint64_add()

//go:wasm-module pantopic/wazero-atomic
//export __atomic_uint64_load
func uint64_load()

//go:wasm-module pantopic/wazero-atomic
//export __atomic_uint64_store
func uint64_store()

// Fix for lint rule `unusedfunc`
var _ = __atomic
