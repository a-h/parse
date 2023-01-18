package parse

type eofParser[T any] struct {
}

func (p eofParser[T]) Parse(in Input) (match T, err error) {
	if _, canAdvance := in.Peek(1); canAdvance {
		err = ErrNotMatched
		return
	}
	return
}

// EOF matches the end of the input.
func EOF[T any]() Parser[T] {
	return eofParser[T]{}
}
