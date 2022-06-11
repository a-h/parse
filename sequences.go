package parse

type sequenceOf2Parser[A, B any] struct {
	A Parser[A]
	B Parser[B]
}

func (p sequenceOf2Parser[A, B]) Parse(in Input) (match SequenceOf2Result[A, B], ok bool, err error) {
	a, ok, err := p.A.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match.A = a
	b, ok, err := p.B.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match.B = b
	ok = true
	return
}

type SequenceOf2Result[A, B any] struct {
	A A
	B B
}

func SequenceOf2[A, B any](a Parser[A], b Parser[B]) Parser[SequenceOf2Result[A, B]] {
	return sequenceOf2Parser[A, B]{
		A: a,
		B: b,
	}
}

type sequenceOf3Parser[A, B, C any] struct {
	Parser[SequenceOf2Result[A, B]]
	C Parser[C]
}

func (p sequenceOf3Parser[A, B, C]) Parse(in Input) (match SequenceOf3Result[A, B, C], ok bool, err error) {
	prefix, ok, err := p.Parser.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match.A = prefix.A
	match.B = prefix.B
	c, ok, err := p.C.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match.C = c
	ok = true
	return
}

type SequenceOf3Result[A, B, C any] struct {
	A A
	B B
	C C
}

func SequenceOf3[A, B, C any](a Parser[A], b Parser[B], c Parser[C]) Parser[SequenceOf3Result[A, B, C]] {
	return sequenceOf3Parser[A, B, C]{
		Parser: SequenceOf2(a, b),
		C:      c,
	}
}
