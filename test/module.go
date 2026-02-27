package main

import (
	"github.com/pantopic/wazero-atomic/sdk-go"
)

const (
	TEST = iota
)

var (
	u64 *atomic.Uint64
)

func main() {
	u64 = atomic.NewUint64(TEST)
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

// Fix for lint rule `unusedfunc`
var _ = testUint64Add
var _ = testUint64Load
var _ = testUint64Store
