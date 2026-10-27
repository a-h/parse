package parse

type orParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

func (p orParser[A, B]) Parse(in *Input) (match Tuple2[Result[A], Result[B]], ok bool, err error) {
	match.A.Value, match.A.OK, err = p.A.Parse(in)
	if err != nil {
		return
	}
	if match.A.OK {
		ok = true
		return
	}

	match.B.Value, match.B.OK, err = p.B.Parse(in)
	if err != nil {
		return
	}
	if match.B.OK {
		ok = true
		return
	}

	return
}

func (p orParser[A, B]) Match(in *Input) (ok bool, err error) {
	ok, err = p.A.Match(in)
	if err != nil || ok {
		return
	}
	return p.B.Match(in)
}

// Or returns a success if either a or b can be parsed.
// If both a and b match, a takes precedence.
func Or[A any, B any](a Parser[A], b Parser[B]) Parser[Tuple2[Result[A], Result[B]]] {
	return orParser[A, B]{
		A: a,
		B: b,
	}
}
