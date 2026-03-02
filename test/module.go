package main

import (
	"github.com/pantopic/wazero-atomic/sdk-go"
)

const (
	ATOMIC_UINT64_SET_1 = iota
	ATOMIC_UINT64_SET_2
)

const (
	ATOMIC_UINT64_ID_A = iota
	ATOMIC_UINT64_ID_B
)

var (
	set1  *atomic.Uint64Set
	u641a *atomic.Uint64
	u641b *atomic.Uint64

	set2  *atomic.Uint64Set
	u642a *atomic.Uint64
	u642b *atomic.Uint64
)

func main() {
	set1 = atomic.NewUint64Set(ATOMIC_UINT64_SET_1)
	u641a = set1.Find(ATOMIC_UINT64_ID_A)
	u641b = set1.Find(ATOMIC_UINT64_ID_B)
	set2 = atomic.NewUint64Set(ATOMIC_UINT64_SET_2)
	u642a = set2.Find(ATOMIC_UINT64_ID_A)
	u642b = set2.Find(ATOMIC_UINT64_ID_B)
}

//export testUint64Add1a
func testUint64Add1a(n uint64) uint64 {
	return u641a.Add(n)
}

//export testUint64Load1a
func testUint64Load1a() uint64 {
	return u641a.Load()
}

//export testUint64Store1a
func testUint64Store1a(n uint64) {
	u641a.Store(n)
}

//export testUint64Del1a
func testUint64Del1a() {
	set1.Del(ATOMIC_UINT64_ID_A)
}

//export testUint64Add1b
func testUint64Add1b(n uint64) uint64 {
	return u641b.Add(n)
}

//export testUint64Load1b
func testUint64Load1b() uint64 {
	return u641b.Load()
}

//export testUint64Store1b
func testUint64Store1b(n uint64) {
	u641b.Store(n)
}

//export testUint64Del1b
func testUint64Del1b() {
	set1.Del(ATOMIC_UINT64_ID_B)
}

//export testUint64Add2a
func testUint64Add2a(n uint64) uint64 {
	return u642a.Add(n)
}

//export testUint64Load2a
func testUint64Load2a() uint64 {
	return u642a.Load()
}

//export testUint64Store2a
func testUint64Store2a(n uint64) {
	u642a.Store(n)
}

//export testUint64Del2a
func testUint64Del2a() {
	set2.Del(ATOMIC_UINT64_ID_A)
}

//export testUint64Add2b
func testUint64Add2b(n uint64) uint64 {
	return u642b.Add(n)
}

//export testUint64Load2b
func testUint64Load2b() uint64 {
	return u642b.Load()
}

//export testUint64Store2b
func testUint64Store2b(n uint64) {
	u642b.Store(n)
}

//export testUint64Del2b
func testUint64Del2b() {
	set2.Del(ATOMIC_UINT64_ID_B)
}

// Fix for lint rule `unusedfunc`
var _ = testUint64Add1a
var _ = testUint64Load1a
var _ = testUint64Store1a
var _ = testUint64Del1a
var _ = testUint64Add1b
var _ = testUint64Load1b
var _ = testUint64Store1b
var _ = testUint64Del1b
var _ = testUint64Add2a
var _ = testUint64Load2a
var _ = testUint64Store2a
var _ = testUint64Del2a
var _ = testUint64Add2b
var _ = testUint64Load2b
var _ = testUint64Store2b
var _ = testUint64Del2b
