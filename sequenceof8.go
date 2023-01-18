package parse

type sequenceOf8Parser[A, B, C, D, E, F, G, H any] struct {
	A Parser[A]
	B Parser[B]
	C Parser[C]
	D Parser[D]
	E Parser[E]
	F Parser[F]
	G Parser[G]
	H Parser[H]
}

func (p sequenceOf8Parser[A, B, C, D, E, F, G, H]) Parse(in Input) (match Tuple8[A, B, C, D, E, F, G, H], err error) {
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
	return
}

func SequenceOf8[A, B, C, D, E, F, G, H any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G], h Parser[H]) Parser[Tuple8[A, B, C, D, E, F, G, H]] {
	return sequenceOf8Parser[A, B, C, D, E, F, G, H]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
	}
}
