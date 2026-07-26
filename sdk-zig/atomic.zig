var _set: u32 = 0;
var _id: u64 = 0;
var _u64: u64 = 0;
var meta = [3]u32{ 0, 0, 0 };

export fn __atomic() u32 {
    meta[0] = @intFromPtr(&_set);
    meta[1] = @intFromPtr(&_id);
    meta[2] = @intFromPtr(&_u64);
    return @intFromPtr(&meta[0]);
}

extern "pantopic/wazero-atomic" fn __atomic_uint64_add() void;
extern "pantopic/wazero-atomic" fn __atomic_uint64_load() void;
extern "pantopic/wazero-atomic" fn __atomic_uint64_store() void;
extern "pantopic/wazero-atomic" fn __atomic_uint64_del() void;

pub const Uint64 = struct {
    set: u32,
    id: u64,

    pub fn add(x: Uint64, delta: u64) u64 {
        _set = x.set;
        _id = x.id;
        _u64 = delta;
        __atomic_uint64_add();
        return _u64;
    }

    pub fn load(x: Uint64) u64 {
        _set = x.set;
        _id = x.id;
        __atomic_uint64_load();
        return _u64;
    }

    pub fn store(x: Uint64, val: u64) void {
        _set = x.set;
        _id = x.id;
        _u64 = val;
        __atomic_uint64_store();
    }
};

pub const Uint64Set = struct {
    set: u32,

    /// init returns a new keyspace of Uint64
    /// set = 0 represents the global set
    pub fn init(set: u32) Uint64Set {
        return .{ .set = set };
    }

    pub fn find(s: Uint64Set, id: u64) Uint64 {
        return .{ .set = s.set, .id = id };
    }

    pub fn add(s: Uint64Set, id: u64, delta: u64) u64 {
        _set = s.set;
        _id = id;
        _u64 = delta;
        __atomic_uint64_add();
        return _u64;
    }

    pub fn load(s: Uint64Set, id: u64) u64 {
        _set = s.set;
        _id = id;
        __atomic_uint64_load();
        return _u64;
    }

    pub fn store(s: Uint64Set, id: u64, val: u64) void {
        _set = s.set;
        _id = id;
        _u64 = val;
        __atomic_uint64_store();
    }

    pub fn del(s: Uint64Set, id: u64) void {
        _set = s.set;
        _id = id;
        __atomic_uint64_del();
    }
};
