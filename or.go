package parse

import "errors"

type orParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

func (p orParser[A, B]) Parse(in Input) (match Tuple2[Match[A], Match[B]], err error) {
	match.A.Value, err = p.A.Parse(in)
	if err != nil && !errors.Is(err, ErrNotMatched) {
		return
	}
	match.A.OK = err == nil
	if match.A.OK {
		return
	}

	match.B.Value, err = p.B.Parse(in)
	if err != nil {
		return
	}
	match.B.OK = err == nil
	if match.B.OK {
		return
	}

	return match, ErrNotMatched
}

// Or returns a success if either a or b can be parsed.
// If both a and b match, a takes precedence.
func Or[A any, B any](a Parser[A], b Parser[B]) Parser[Tuple2[Match[A], Match[B]]] {
	return orParser[A, B]{
		A: a,
		B: b,
	}
}
