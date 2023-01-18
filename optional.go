package parse

import "errors"

type optionalParser[T any] struct {
	Parser      Parser[T]
	Insensitive bool
}

func (p optionalParser[T]) Parse(in Input) (match Match[T], err error) {
	match.Value, err = p.Parser.Parse(in)
	if errors.Is(err, ErrNotMatched) {
		match.OK = false
		return match, nil
	}
	if err != nil {
		return
	}
	match.OK = true
	return
}

type Match[T any] struct {
	Value T
	OK    bool
}

// Optional converts the given parser into an optional parser.
func Optional[T any](parser Parser[T]) Parser[Match[T]] {
	return optionalParser[T]{
		Parser: parser,
	}
}
