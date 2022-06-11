package parse

import "strings"

type StringParser struct {
	Match       string
	Insensitive bool
}

func (p StringParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(len(p.Match))
	if !ok {
		return
	}
	if p.Insensitive {
		ok = strings.EqualFold(p.Match, match)
	} else {
		ok = p.Match == match
	}
	if !ok {
		match = ""
		return
	}
	in.Chomp(len(p.Match))
	return
}

func String(s string) Parser[string] {
	return StringParser{
		Match: s,
	}
}

func StringInsensitive(s string) Parser[string] {
	return StringParser{
		Match:       s,
		Insensitive: true,
	}
}
