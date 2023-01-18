package parse

type sequenceOf4Parser[A, B, C, D any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
	D Parser[D]
}

func (p sequenceOf4Parser[A, B, C, D]) Parse(in Input) (match Tuple4[A, B, C, D], err error) {
	match.A, err = p.A.Parse(in)
	if err != nil {
		return
	}
	match.B, err = p.B.Parse(in)
	if err != nil {
		return
	}
	match.C, err = p.C.Parse(in)
	if err != nil {
		return
	}
	match.D, err = p.D.Parse(in)
	if err != nil {
		return
	}
	return
}

func SequenceOf4[A, B, C, D any](a Parser[A], b Parser[B], c Parser[C], d Parser[D]) Parser[Tuple4[A, B, C, D]] {
	return sequenceOf4Parser[A, B, C, D]{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}
