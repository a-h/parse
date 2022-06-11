package parse

import (
	"strings"
)

type RuneInParser struct {
	Runes string
}

func (p RuneInParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(1)
	if !ok {
		return
	}
	ok = strings.Contains(p.Runes, match)
	if !ok {
		return
	}
	in.Chomp(1)
	return
}

func RuneIn(s string) Parser[string] {
	return RuneInParser{
		Runes: s,
	}
}
