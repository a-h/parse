package parse

type stringUntilParser struct {
	Delimiter Matcher
	AllowEOF  bool
}

func (p stringUntilParser) Parse(in *Input) (match string, ok bool, err error) {
	start := in.Index()
	ok, err = p.Match(in)
	if err != nil || !ok {
		return
	}
	end := in.Index()
	in.Seek(start)
	match, ok = in.Take(end - start)
	return
}

func (p stringUntilParser) Match(in *Input) (ok bool, err error) {
	start := in.Index()
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
		_, chompOK := in.Take(1)
		if !chompOK {
			if p.AllowEOF {
				return true, nil
			}
			in.Seek(start)
			return false, nil
		}
	}
}

// StringUntil matches until the delimiter is reached.
func StringUntil(delimiter Matcher) Parser[string] {
	return stringUntilParser{
		Delimiter: delimiter,
	}
}

// StringUntilEOF matches until the delimiter or the end of the file is reached.
func StringUntilEOF(delimiter Matcher) Parser[string] {
	return stringUntilParser{
		Delimiter: delimiter,
		AllowEOF:  true,
	}
}
