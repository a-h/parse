package parse

type eofParser[T any] struct {
}

func (p eofParser[T]) Parse(in *Input) (match T, ok bool, err error) {
	_, canAdvance := in.Peek(1)
	ok = !canAdvance
	return
}

func (p eofParser[T]) Match(in *Input) (ok bool, err error) {
	_, ok, err = p.Parse(in)
	return
}

// EOF matches the end of the input.
func EOF[T any]() Parser[T] {
	return eofParser[T]{}
}
