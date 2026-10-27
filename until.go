package parse

type untilParser[T any] struct {
	Parser    Parser[T]
	Delimiter Matcher
	AllowEOF  bool
}

func (p untilParser[T]) Parse(in *Input) (match []T, ok bool, err error) {
	start := in.Index()
	if _, ok = in.Peek(1); !ok && p.AllowEOF {
		ok = true
		return
	}
	var m T
	m, ok, err = p.Parser.Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	match = append(match, m)
	for {
		beforeDelimiter := in.Index()
		ok, err = p.Delimiter.Match(in)
		if err != nil {
			in.Seek(start)
			return
		}
		if ok {
			in.Seek(beforeDelimiter)
			return
		}
		if _, ok = in.Peek(1); !ok && p.AllowEOF {
			ok = true
			return
		}
		var m T
		m, ok, err = p.Parser.Parse(in)
		if err != nil {
			in.Seek(start)
			return
		}
		if !ok {
			in.Seek(start)
			return
		}
		match = append(match, m)
	}
}

func (p untilParser[T]) Match(in *Input) (ok bool, err error) {
	start := in.Index()
	if _, ok = in.Peek(1); !ok && p.AllowEOF {
		return true, nil
	}
	ok, err = p.Parser.Match(in)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	for {
		beforeDelimiter := in.Index()
		ok, err = p.Delimiter.Match(in)
		if err != nil {
			in.Seek(start)
			return
		}
		if ok {
			in.Seek(beforeDelimiter)
			return
		}
		if _, ok = in.Peek(1); !ok && p.AllowEOF {
			return true, nil
		}
		ok, err = p.Parser.Match(in)
		if err != nil {
			in.Seek(start)
			return
		}
		if !ok {
			in.Seek(start)
			return
		}
	}
}

// Until matches until the delimiter is reached.
func Until[T any](parser Parser[T], delimiter Matcher) Parser[[]T] {
	return untilParser[T]{
		Parser:    parser,
		Delimiter: delimiter,
	}
}

// UntilEOF matches until the delimiter or the end of the file is reached.
func UntilEOF[T any](parser Parser[T], delimiter Matcher) Parser[[]T] {
	return untilParser[T]{
		Parser:    parser,
		Delimiter: delimiter,
		AllowEOF:  true,
	}
}
