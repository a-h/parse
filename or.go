package parse

type orParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

func (p orParser[A, B]) Parse(in Input) (match Tuple2[Match[A], Match[B]], ok bool, err error) {
	a, ok, err := p.A.Parse(in)
	if err != nil {
		return
	}
	if ok {
		match = Tuple2[Match[A], Match[B]]{
			A: Match[A]{
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
		match = Tuple2[Match[A], Match[B]]{
			B: Match[B]{
				Value: b,
				OK:    true,
			},
		}
		return
	}
	return
}

// Or returns a success if either a or b can be parsed.
// If both a and b match, a takes precedence.
func Or[A any, B any](a Parser[A], b Parser[B]) Parser[Tuple2[Match[A], Match[B]]] {
	return orParser[A, B]{
		A: a,
		B: b,
	}
}
