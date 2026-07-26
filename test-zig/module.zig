const atomic = @import("atomic");

const ATOMIC_UINT64_SET_1 = 0;
const ATOMIC_UINT64_SET_2 = 1;

const ATOMIC_UINT64_ID_A = 0;
const ATOMIC_UINT64_ID_B = 1;

const set1 = atomic.Uint64Set.init(ATOMIC_UINT64_SET_1);
const u641a = set1.find(ATOMIC_UINT64_ID_A);
const u641b = set1.find(ATOMIC_UINT64_ID_B);

const set2 = atomic.Uint64Set.init(ATOMIC_UINT64_SET_2);
const u642a = set2.find(ATOMIC_UINT64_ID_A);
const u642b = set2.find(ATOMIC_UINT64_ID_B);

export fn testUint64Add1a(n: u64) u64 {
    return u641a.add(n);
}

export fn testUint64Load1a() u64 {
    return u641a.load();
}

export fn testUint64Store1a(n: u64) void {
    u641a.store(n);
}

export fn testUint64Del1a() void {
    set1.del(ATOMIC_UINT64_ID_A);
}

export fn testUint64Add1b(n: u64) u64 {
    return u641b.add(n);
}

export fn testUint64Load1b() u64 {
    return u641b.load();
}

export fn testUint64Store1b(n: u64) void {
    u641b.store(n);
}

export fn testUint64Del1b() void {
    set1.del(ATOMIC_UINT64_ID_B);
}

export fn testUint64Add2a(n: u64) u64 {
    return u642a.add(n);
}

export fn testUint64Load2a() u64 {
    return u642a.load();
}

export fn testUint64Store2a(n: u64) void {
    u642a.store(n);
}

export fn testUint64Del2a() void {
    set2.del(ATOMIC_UINT64_ID_A);
}

export fn testUint64Add2b(n: u64) u64 {
    return u642b.add(n);
}

export fn testUint64Load2b() u64 {
    return u642b.load();
}

export fn testUint64Store2b(n: u64) void {
    u642b.store(n);
}

export fn testUint64Del2b() void {
    set2.del(ATOMIC_UINT64_ID_B);
}
