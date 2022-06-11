package parse

import (
	"strings"
	"unicode"
)

func Rune(r rune) Parser[string] {
	return StringParser{
		Match: string(r),
	}
}

type RuneWhereParser struct {
	F func(r rune) bool
}

func (p RuneWhereParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(1)
	if !ok {
		return
	}
	ok = p.F(rune(match[0]))
	if !ok {
		return
	}
	in.Chomp(1)
	return
}

func RuneWhere(predicate func(r rune) bool) Parser[string] {
	return RuneWhereParser{
		F: predicate,
	}
}

func RuneIn(s string) Parser[string] {
	return RuneWhere(func(r rune) bool { return strings.Contains(s, string(r)) })
}

func RuneNotIn(s string) Parser[string] {
	return RuneWhere(func(r rune) bool { return !strings.Contains(s, string(r)) })
}

func RuneInRanges(ranges ...*unicode.RangeTable) Parser[string] {
	return RuneWhere(func(r rune) bool { return unicode.IsOneOf(ranges, r) })
}

// AnyRune matches any single rune.
var AnyRune = RuneWhere(func(r rune) bool { return true })
