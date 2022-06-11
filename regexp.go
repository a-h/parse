package parse

import (
	"regexp"
)

type RegexpParser struct {
	Expression *regexp.Regexp
}

func (p RegexpParser) Parse(in Input) (match string, ok bool, err error) {
	remainder, ok := in.Peek(-1)
	if !ok {
		return
	}
	startAndEndIndex := p.Expression.FindStringIndex(remainder)
	ok = startAndEndIndex != nil && startAndEndIndex[0] == 0
	if !ok {
		return
	}
	match = remainder[startAndEndIndex[0]:startAndEndIndex[1]]
	in.Chomp(len(match))
	return
}

func Regexp(s string) (p Parser[string], err error) {
	r, err := regexp.Compile(s)
	if err != nil {
		return
	}
	p = RegexpParser{
		Expression: r,
	}
	return
}

func MustRegexp(s string) (p Parser[string]) {
	p, err := Regexp(s)
	if err != nil {
		panic(err)
	}
	return
}
