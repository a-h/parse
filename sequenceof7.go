package parse

type sequenceOf7Parser[A, B, C, D, E, F, G any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
	D Parser[D]
	E Parser[E]
	F Parser[F]
	G Parser[G]
}

func (p sequenceOf7Parser[A, B, C, D, E, F, G]) Parse(in Input) (match Tuple7[A, B, C, D, E, F, G], err error) {
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
	match.G, err = p.G.Parse(in)
	if err != nil {
		return
	}
	return
}

func SequenceOf7[A, B, C, D, E, F, G any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G]) Parser[Tuple7[A, B, C, D, E, F, G]] {
	return sequenceOf7Parser[A, B, C, D, E, F, G]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
	}
}
