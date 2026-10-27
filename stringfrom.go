package parse

type stringFromParser struct {
	Matchers []Matcher
}

func (p stringFromParser) Parse(in *Input) (match string, ok bool, err error) {
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

func (p stringFromParser) Match(in *Input) (ok bool, err error) {
	start := in.Index()
	for _, matcher := range p.Matchers {
		ok, err = matcher.Match(in)
		if err != nil {
			return
		}
		if !ok {
			in.Seek(start)
			return
		}
	}
	return
}

// StringFrom returns the string range captured by the given matchers.
func StringFrom(matchers ...Matcher) Parser[string] {
	return stringFromParser{
		Matchers: matchers,
	}
}
