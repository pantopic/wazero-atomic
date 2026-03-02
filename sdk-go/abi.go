package atomic

import (
	"unsafe"
)

var (
	_set uint32
	_id  uint64
	_u64 uint64
	meta = make([]uint32, 3)
)

//export __atomic
func __atomic() (res uint32) {
	meta[0] = uint32(uintptr(unsafe.Pointer(&_set)))
	meta[1] = uint32(uintptr(unsafe.Pointer(&_id)))
	meta[2] = uint32(uintptr(unsafe.Pointer(&_u64)))
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

//go:wasm-module pantopic/wazero-atomic
//export __atomic_uint64_del
func uint64_del()

// Fix for lint rule `unusedfunc`
var _ = __atomic
