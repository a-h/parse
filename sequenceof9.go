package parse

type sequenceOf9Parser[A, B, C, D, E, F, G, H, I any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
	D Parser[D]
	E Parser[E]
	F Parser[F]
	G Parser[G]
	H Parser[H]
	I Parser[I]
}

func (p sequenceOf9Parser[A, B, C, D, E, F, G, H, I]) Parse(in Input) (match Tuple9[A, B, C, D, E, F, G, H, I], err error) {
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
	match.H, err = p.H.Parse(in)
	if err != nil {
		return
	}
	match.I, err = p.I.Parse(in)
	if err != nil {
		return
	}
	return
}

func SequenceOf9[A, B, C, D, E, F, G, H, I any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G], h Parser[H], i Parser[I]) Parser[Tuple9[A, B, C, D, E, F, G, H, I]] {
	return sequenceOf9Parser[A, B, C, D, E, F, G, H, I]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
	}
}
