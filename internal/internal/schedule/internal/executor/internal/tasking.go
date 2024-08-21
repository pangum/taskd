package internal

type Tasking struct {
	id     uint64
	target uint64
	data   any
	times  uint32
}

func NewTasking(id uint64, target uint64, data any, times uint32) *Tasking {
	return &Tasking{
		id:     id,
		target: target,
		data:   data,
		times:  times,
	}
}

func (t *Tasking) Id() uint64 {
	return t.id
}

func (t *Tasking) Target() uint64 {
	return t.target
}

func (t *Tasking) Data() any {
	return t.data
}

func (t *Tasking) Times() uint32 {
	return t.times
}
