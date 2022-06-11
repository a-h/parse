package parse

type OrParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

type OrResult[A any, B any] struct {
	A OptionalResult[A]
	B OptionalResult[B]
}

func (p OrParser[A, B]) Parse(in Input) (match OrResult[A, B], ok bool, err error) {
	a, ok, err := p.A.Parse(in)
	if err != nil {
		return
	}
	if ok {
		match = OrResult[A, B]{
			A: OptionalResult[A]{
				Value: a,
				OK:    true,
			},
		}
		return
	}
	b, ok, err := p.B.Parse(in)
	if err != nil {
		return
	}
	if ok {
		match = OrResult[A, B]{
			B: OptionalResult[B]{
				Value: b,
				OK:    true,
			},
		}
		return
	}
	return
}

func Or[A any, B any](a Parser[A], b Parser[B]) OrParser[A, B] {
	return OrParser[A, B]{
		A: a,
		B: b,
	}
}
