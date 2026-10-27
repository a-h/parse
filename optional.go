package parse

type optionalParser[T any] struct {
	Parser      Parser[T]
	Insensitive bool
}

func (p optionalParser[T]) Parse(in *Input) (match Result[T], ok bool, err error) {
	match.Value, match.OK, err = p.Parser.Parse(in)
	if err != nil {
		return
	}
	return match, true, nil
}

func (p optionalParser[T]) Match(in *Input) (ok bool, err error) {
	_, err = p.Parser.Match(in)
	if err != nil {
		return
	}
	return true, nil
}

// Result is the value of a parser that can succeed without matching, such as
// Optional. OK reports whether the parser matched.
type Result[T any] struct {
	Value T
	OK    bool
}

// Optional converts the given parser into an optional parser.
func Optional[T any](parser Parser[T]) Parser[Result[T]] {
	return optionalParser[T]{
		Parser: parser,
	}
}
