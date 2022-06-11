package parse

import "io"

type SequenceOf2Parser[A, B any] struct {
	A Parser[A]
	B Parser[B]
}

type SequenceOf2Result[A, B any] struct {
	A A
	B B
}

func (p SequenceOf2Parser[A, B]) Parse(in Input) (match SequenceOf2Result[A, B], ok bool, err error) {
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

func SequenceOf2[A, B any](a Parser[A], b Parser[B]) SequenceOf2Parser[A, B] {
	return SequenceOf2Parser[A, B]{
		A: a,
		B: b,
	}
}

type SequenceOf3Parser[A, B, C any] struct {
	SequenceOf2Parser[A, B]
	C Parser[C]
}

type SequenceOf3Result[A, B, C any] struct {
	A A
	B B
	C C
}

func (p SequenceOf3Parser[A, B, C]) Parse(in Input) (match SequenceOf3Result[A, B, C], ok bool, err error) {
	prefix, ok, err := p.SequenceOf2Parser.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match.A = prefix.A
	match.B = prefix.B
	c, ok, err := p.C.Parse(in)
	if err != nil && err != io.EOF {
		return
	}
	if !ok {
		return
	}
	match.C = c
	ok = true
	return
}

func SequenceOf3[A, B, C any](a Parser[A], b Parser[B], c Parser[C]) SequenceOf3Parser[A, B, C] {
	return SequenceOf3Parser[A, B, C]{
		SequenceOf2Parser: SequenceOf2(a, b),
		C:                 c,
	}
}
