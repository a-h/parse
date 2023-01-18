package parse

type sequenceOf6Parser[A, B, C, D, E, F any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
	D Parser[D]
	E Parser[E]
	F Parser[F]
}

func (p sequenceOf6Parser[A, B, C, D, E, F]) Parse(in Input) (match Tuple6[A, B, C, D, E, F], err error) {
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
	match.E, err = p.E.Parse(in)
	if err != nil {
		return
	}
	match.F, err = p.F.Parse(in)
	if err != nil {
		return
	}
	return
}

func SequenceOf6[A, B, C, D, E, F any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F]) Parser[Tuple6[A, B, C, D, E, F]] {
	return sequenceOf6Parser[A, B, C, D, E, F]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}
