package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

type prefixParser struct {
	prefix string
}

func (p prefixParser) Parse(in *parse.Input) (match string, ok bool, err error) {
	match, ok = in.Peek(len(p.prefix))
	if !ok || match != p.prefix {
		return "", false, nil
	}
	in.Take(len(p.prefix))
	return
}

func (p prefixParser) Match(in *parse.Input) (ok bool, err error) {
	return parse.Match(p, in)
}

func TestParsersDoNotAllocate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		parser parse.Parser[string]
	}{
		{name: "Whitespace", input: "  \t\n  x", parser: parse.Whitespace},
		{name: "OptionalWhitespace", input: "  \t\n  x", parser: parse.OptionalWhitespace},
		{name: "StringFrom OneOrMore", input: "abc ", parser: parse.StringFrom(parse.OneOrMore(parse.Letter))},
		{name: "StringFrom All", input: "</div>", parser: parse.StringFrom(parse.All(parse.String("</"), parse.String("div"), parse.Rune('>')))},
		{name: "StringFrom Times", input: "2026", parser: parse.StringFrom(parse.Times(4, parse.ZeroToNine))},
		{name: "StringUntil", input: "abc-->", parser: parse.StringUntil(parse.String("-->"))},
		{name: "StringFrom Until", input: "abc;", parser: parse.StringFrom(parse.Until(parse.Letter, parse.Rune(';')))},
		{name: "StringFrom a parser that uses Match", input: "abc", parser: parse.StringFrom(prefixParser{prefix: "abc"})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := parse.NewInput([]byte(tt.input))
			allocs := testing.AllocsPerRun(100, func() {
				in.Seek(0)
				_, ok, err := tt.parser.Parse(in)
				if err != nil || !ok {
					t.Fatalf("failed to parse: ok=%v, err=%v", ok, err)
				}
			})
			if allocs != 0 {
				t.Errorf("got %v allocations, expected 0", allocs)
			}
		})
	}
}

func TestASCIIRuneConstructionDoesNotAllocate(t *testing.T) {
	var p parse.Parser[string]
	allocs := testing.AllocsPerRun(100, func() {
		p = parse.Rune('>')
	})
	if allocs != 0 {
		t.Errorf("got %v allocations, expected 0", allocs)
	}
	if p == nil {
		t.Fatal("expected a parser")
	}
}
