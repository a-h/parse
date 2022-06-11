package parse

type StringFromParser[T any] struct {
	Parsers []Parser[T]
}

func (p StringFromParser[T]) Parse(in Input) (match string, ok bool, err error) {
	start := in.Index()
	for _, parser := range p.Parsers {
		_, ok, err = parser.Parse(in)
		if err != nil {
			return
		}
		if !ok {
			in.Seek(start)
			return
		}
	}
	end := in.Index()
	in.Seek(start)
	match, ok = in.Chomp(end - start)
	return
}

func StringFrom[T any](parsers ...Parser[T]) Parser[string] {
	return StringFromParser[T]{
		Parsers: parsers,
	}
}
