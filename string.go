package parse

type StringParser struct {
	Match string
}

func (sp StringParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(len(sp.Match))
	if !ok {
		return
	}
	ok = sp.Match == match
	if !ok {
		return
	}
	in.Chomp(len(sp.Match))
	return
}

func String(s string) Parser[string] {
	return StringParser{
		Match: s,
	}
}
