package parse

type allParser[T any] struct {
	Parsers []Parser[T]
}

func (p allParser[T]) Parse(in Input) (match []T, err error) {
	for _, parser := range p.Parsers {
		var m T
		m, err = parser.Parse(in)
		if err != nil {
			return
		}
		match = append(match, m)
	}
	return
}

// All parses all of the parsers in the list in sequence and combines the result.
func All[T any](parsers ...Parser[T]) Parser[[]T] {
	return allParser[T]{
		Parsers: parsers,
	}
}
