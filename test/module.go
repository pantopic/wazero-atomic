package main

import (
	"github.com/pantopic/wazero-atomic/sdk-go"
)

const (
	ATOMIC_UINT64_ID_TEST = iota
	ATOMIC_UINT64_ID_TEST_2
)

var (
	u64  *atomic.Uint64
	u64b *atomic.Uint64
)

func main() {
	u64 = atomic.NewUint64(ATOMIC_UINT64_ID_TEST)
	u64b = atomic.NewUint64(ATOMIC_UINT64_ID_TEST_2)
}

//export testUint64Add
func testUint64Add(n uint64) uint64 {
	return u64.Add(n)
}

//export testUint64Load
func testUint64Load() uint64 {
	return u64.Load()
}

//export testUint64Store
func testUint64Store(n uint64) {
	u64.Store(n)
}

//export testUint64Add2
func testUint64Add2(n uint64) uint64 {
	return u64b.Add(n)
}

//export testUint64Load2
func testUint64Load2() uint64 {
	return u64b.Load()
}

//export testUint64Store2
func testUint64Store2(n uint64) {
	u64b.Store(n)
}

// Fix for lint rule `unusedfunc`
var _ = testUint64Add
var _ = testUint64Load
var _ = testUint64Store
var _ = testUint64Add2
var _ = testUint64Load2
var _ = testUint64Store2
