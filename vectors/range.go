package vectors

type Range struct {
	Start, End int
}

func (r Range) Total() int {
	return r.End - r.Start
}
