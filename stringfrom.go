package parse

import "errors"

type stringFromParser[T any] struct {
	Parsers []Parser[T]
}

func (p stringFromParser[T]) Parse(in Input) (match string, err error) {
	start := in.Index()
	for _, parser := range p.Parsers {
		_, err = parser.Parse(in)
		if errors.Is(err, ErrNotMatched) {
			in.Seek(start)
			return
		}
		if err != nil {
			return
		}
	}
	end := in.Index()
	in.Seek(start)
	match, _ = in.Take(end - start)
	return
}

// StringFrom returns the string range captured by the given parsers.
func StringFrom[T any](parsers ...Parser[T]) Parser[string] {
	return stringFromParser[T]{
		Parsers: parsers,
	}
}
