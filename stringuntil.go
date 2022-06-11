package parse

type StringUntilParser[T any] struct {
	Delimiter Parser[T]
	AllowEOF  bool
}

func (p StringUntilParser[T]) Parse(in Input) (match string, ok bool, err error) {
	start := in.Index()
	for {
		beforeDelimiter := in.Index()
		_, ok, err = p.Delimiter.Parse(in)
		if err != nil {
			return
		}
		if ok {
			in.Seek(beforeDelimiter)
			break
		}
		_, chompOK := in.Chomp(1)
		if !chompOK {
			if p.AllowEOF {
				break
			}
			return "", false, nil
		}
	}
	end := in.Index()
	in.Seek(start)
	match, ok = in.Chomp(end - start)
	return
}

func StringUntil[T any](delimiter Parser[T]) Parser[string] {
	return StringUntilParser[T]{
		Delimiter: delimiter,
	}
}

func StringUntilEOF[T any](delimiter Parser[T]) Parser[string] {
	return StringUntilParser[T]{
		Delimiter: delimiter,
		AllowEOF:  true,
	}
}
