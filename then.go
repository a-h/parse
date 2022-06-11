package parse

import "io"

type ThenParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

type ThenResult[A any, B any] struct {
	A A
	B B
}

func (p ThenParser[A, B]) Parse(in Input) (match ThenResult[A, B], ok bool, err error) {
	a, ok, err := p.A.Parse(in)
	if err != nil && err != io.EOF {
		return
	}
	if !ok {
		return
	}
	match.A = a
	b, ok, err := p.B.Parse(in)
	if err != nil && err != io.EOF {
		return
	}
	if !ok {
		return
	}
	match.B = b
	ok = true
	return
}

func Then[A any, B any](a Parser[A], b Parser[B]) ThenParser[A, B] {
	return ThenParser[A, B]{
		A: a,
		B: b,
	}
}
