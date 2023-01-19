package parse

import "errors"

type anyParser[T any] struct {
	Parsers []Parser[T]
}

func (p anyParser[T]) Parse(in Input) (match T, err error) {
	for _, parser := range p.Parsers {
		match, err = parser.Parse(in)
		if err != nil && !errors.Is(err, ErrNotMatched) {
			return match, err
		}
		if errors.Is(err, ErrNotMatched) {
			continue
		}
		return match, err
	}
	return
}

// Any parses any one of the parsers in the list.
func Any[T any](parsers ...Parser[T]) Parser[T] {
	return anyParser[T]{
		Parsers: parsers,
	}
}
