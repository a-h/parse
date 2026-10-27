package parse

import "strings"

type stringParser struct {
	Value       string
	Insensitive bool
}

func (p stringParser) Parse(in *Input) (match string, ok bool, err error) {
	match, ok = in.Peek(len(p.Value))
	if !ok {
		return
	}
	if p.Insensitive {
		ok = strings.EqualFold(p.Value, match)
	} else {
		ok = p.Value == match
	}
	if !ok {
		return "", false, nil
	}
	in.Take(len(p.Value))
	return
}

func (p stringParser) Match(in *Input) (ok bool, err error) {
	_, ok, err = p.Parse(in)
	return
}

// String matches a given string constant.
func String(s string) Parser[string] {
	return stringParser{
		Value: s,
	}
}

// StringInsensitive matches a given string constant using Unicode case folding.
func StringInsensitive(s string) Parser[string] {
	return stringParser{
		Value:       s,
		Insensitive: true,
	}
}
