package parse

type OptionalParser[T any] struct {
	Parser      Parser[T]
	Insensitive bool
}

type OptionalResult[T any] struct {
	Value T
	OK    bool
}

func (p OptionalParser[T]) Parse(in Input) (match OptionalResult[T], ok bool, err error) {
	var item T
	item, ok, err = p.Parser.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return match, true, nil
	}
	match = OptionalResult[T]{
		Value: item,
		OK:    ok,
	}
	return match, true, nil
}

func Optional[T any](parser Parser[T]) Parser[OptionalResult[T]] {
	return OptionalParser[T]{
		Parser: parser,
	}
}
