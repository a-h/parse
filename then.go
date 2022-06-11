package parse

func Then[A any, B any](a Parser[A], b Parser[B]) Parser[SequenceOf2Result[A, B]] {
	return SequenceOf2Parser[A, B]{
		A: a,
		B: b,
	}
}
