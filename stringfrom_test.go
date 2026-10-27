package parse_test

import (
	"errors"
	"testing"

	"github.com/a-h/parse"
)

func TestStringFrom(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.StringFrom(parse.String("ABC"), parse.String("123")),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "matches",
			input:         "ABCDEF",
			parser:        parse.StringFrom(parse.MustRegexp("."), parse.MustRegexp("BC"), parse.String("DEF")),
			expectedMatch: "ABCDEF",
			expectedOK:    true,
		},
		{
			name:          "parsers of different types can be combined",
			input:         "2026-10-27",
			parser:        parse.StringFrom(parse.Times(4, parse.ZeroToNine), parse.Rune('-'), parse.Optional(parse.Times(2, parse.ZeroToNine)), parse.Rune('-'), parse.Times(2, parse.ZeroToNine)),
			expectedMatch: "2026-10-27",
			expectedOK:    true,
		},
		{
			name:          "Until can be captured",
			input:         "abc;",
			parser:        parse.StringFrom(parse.Until(parse.Letter, parse.Rune(';')), parse.Rune(';')),
			expectedMatch: "abc;",
			expectedOK:    true,
		},
		{
			name:          "Times can be captured",
			input:         "AAAB",
			parser:        parse.StringFrom(parse.Times(3, parse.Rune('A'))),
			expectedMatch: "AAA",
			expectedOK:    true,
		},
		{
			name:       "Times rolls back when there are too few matches",
			input:      "AAB",
			parser:     parse.StringFrom(parse.Times(3, parse.Rune('A'))),
			expectedOK: false,
		},
		{
			name:          "Repeat stops at the maximum",
			input:         "AAAA",
			parser:        parse.StringFrom(parse.Repeat(1, 2, parse.Rune('A'))),
			expectedMatch: "AA",
			expectedOK:    true,
		},
		{
			name:          "ZeroOrMore can match nothing",
			input:         "B",
			parser:        parse.StringFrom(parse.ZeroOrMore(parse.Rune('A'))),
			expectedMatch: "",
			expectedOK:    true,
		},
		{
			name:          "All can be captured",
			input:         "ABC",
			parser:        parse.StringFrom(parse.All(parse.Rune('A'), parse.Rune('B'))),
			expectedMatch: "AB",
			expectedOK:    true,
		},
		{
			name:       "All rolls back on a partial match",
			input:      "AC",
			parser:     parse.StringFrom(parse.All(parse.Rune('A'), parse.Rune('B'))),
			expectedOK: false,
		},
		{
			name:          "Any can be captured",
			input:         "B",
			parser:        parse.StringFrom(parse.Any(parse.Rune('A'), parse.Rune('B'))),
			expectedMatch: "B",
			expectedOK:    true,
		},
		{
			name:          "Optional can match nothing",
			input:         "B",
			parser:        parse.StringFrom(parse.Optional(parse.Rune('A')), parse.Optional(parse.Rune('B'))),
			expectedMatch: "B",
			expectedOK:    true,
		},
		{
			name:          "Or can be captured",
			input:         "12",
			parser:        parse.StringFrom(parse.Or(parse.Rune('A'), parse.ZeroToNine), parse.Or(parse.Rune('2'), parse.ZeroToNine)),
			expectedMatch: "12",
			expectedOK:    true,
		},
		{
			name:          "nested slice parsers can be captured",
			input:         "AB AB AB.",
			parser:        parse.StringFrom(parse.OneOrMore(parse.All(parse.String("AB"), parse.StringFrom(parse.Optional(parse.Rune(' ')))))),
			expectedMatch: "AB AB AB",
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}

func TestStringFromReturnsErrorsFromMatchers(t *testing.T) {
	errExpected := errors.New("expected error")
	failing := parse.Func(func(in *parse.Input) (match string, ok bool, err error) {
		return "", false, errExpected
	})
	tests := []ParserTest[string]{
		{name: "Times", input: "A", parser: parse.StringFrom(parse.Times(1, failing)), expectedErr: errExpected},
		{name: "All", input: "A", parser: parse.StringFrom(parse.All(failing)), expectedErr: errExpected},
		{name: "Any", input: "A", parser: parse.StringFrom(parse.Any(failing)), expectedErr: errExpected},
		{name: "Optional", input: "A", parser: parse.StringFrom(parse.Optional(failing)), expectedErr: errExpected},
		{name: "Or", input: "A", parser: parse.StringFrom(parse.Or(parse.Rune('B'), failing)), expectedErr: errExpected},
	}
	RunParserTests(t, tests)
}
