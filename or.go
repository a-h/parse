package parse

type orParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

func (p orParser[A, B]) Parse(in Input) (match OrResult[A, B], ok bool, err error) {
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

type OrResult[A any, B any] struct {
	A OptionalResult[A]
	B OptionalResult[B]
}

// Or returns a success if either a or b can be parsed.
// If both a and b match, a takes precedence.
func Or[A any, B any](a Parser[A], b Parser[B]) Parser[OrResult[A, B]] {
	return orParser[A, B]{
		A: a,
		B: b,
	}
}
