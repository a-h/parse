package parse

type sequenceOf3Parser[A, B, C any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
}

func (p sequenceOf3Parser[A, B, C]) Parse(in Input) (match Tuple3[A, B, C], err error) {
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
	return
}

func SequenceOf3[A, B, C any](a Parser[A], b Parser[B], c Parser[C]) Parser[Tuple3[A, B, C]] {
	return sequenceOf3Parser[A, B, C]{
		A: a,
		B: b,
		C: c,
	}
}
