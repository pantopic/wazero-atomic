package atomic

type Uint64 struct {
	id uint32
}

func NewUint64(id uint32) *Uint64 {
	return &Uint64{uint32(id)}
}

func (x *Uint64) Add(delta uint64) (new uint64) {
	id = x.id
	u64 = delta
	add()
	return u64
}

func (x *Uint64) Load() uint64 {
	load()
	return u64
}

func (x *Uint64) Store(val uint64) {
	id = x.id
	u64 = val
	store()
}
