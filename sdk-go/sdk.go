package atomic

type Uint64 struct {
	set uint32
	id  uint64
}

func (x *Uint64) Add(delta uint64) (new uint64) {
	_set = x.set
	_id = x.id
	_u64 = delta
	uint64_add()
	return _u64
}

func (x *Uint64) Load() uint64 {
	_set = x.set
	_id = x.id
	uint64_load()
	return _u64
}

func (x *Uint64) Store(val uint64) {
	_set = x.set
	_id = x.id
	_u64 = val
	uint64_store()
}

type Uint64Set struct {
	set uint32
}

// NewUint64Set returns a new keyspace of *atomic.Uint64
// [set] = 0 represents the global set
func NewUint64Set(set uint32) *Uint64Set {
	return &Uint64Set{set}
}

func (s *Uint64Set) Find(id uint64) *Uint64 {
	return &Uint64{s.set, id}
}

func (x *Uint64Set) Add(id uint64, delta uint64) (new uint64) {
	_set = x.set
	_id = id
	_u64 = delta
	uint64_add()
	return _u64
}

func (x *Uint64Set) Load(id uint64) uint64 {
	_set = x.set
	_id = id
	uint64_load()
	return _u64
}

func (x *Uint64Set) Store(id uint64, val uint64) {
	_set = x.set
	_id = id
	_u64 = val
	uint64_store()
}

func (x *Uint64Set) Del(id uint64) {
	_set = x.set
	_id = id
	uint64_del()
}
