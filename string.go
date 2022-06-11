package parse

import "strings"

type StringParser struct {
	Match       string
	Insensitive bool
}

func (sp StringParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(len(sp.Match))
	if !ok {
		return
	}
	if sp.Insensitive {
		ok = strings.EqualFold(sp.Match, match)
	} else {
		ok = sp.Match == match
	}
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

func StringInsensitive(s string) Parser[string] {
	return StringParser{
		Match:       s,
		Insensitive: true,
	}
}
