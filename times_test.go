package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestTimes(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			name:        "Times: no match",
			input:       "ABCDEF",
			parser:      parse.Times(2, parse.String("A")),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "Times: matches",
			input:         "AAAA",
			parser:        parse.Times(3, parse.String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedErr:   nil,
		},
		{
			name:          "Repeat: must be at least 1, and take up to 5",
			input:         "AAAA",
			parser:        parse.Repeat(1, 5, parse.String("A")),
			expectedMatch: []string{"A", "A", "A", "A"},
			expectedErr:   nil,
		},
		{
			name:        "Repeat: min of 4, max of 5 - no match",
			input:       "AAA",
			parser:      parse.Repeat(4, 5, parse.String("A")),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "Repeat: min of 0, max of 2 - matches",
			input:         "AAA",
			parser:        parse.Repeat(0, 2, parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedErr:   nil,
		},
		{
			name:          "AtMost: success",
			input:         "AAA",
			parser:        parse.AtMost(2, parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedErr:   nil,
		},
		{
			name:          "AtLeast: success",
			input:         "AAA",
			parser:        parse.AtLeast(2, parse.String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedErr:   nil,
		},
		{
			name:          "ZeroOrMore: nothing to get",
			input:         "BB",
			parser:        parse.ZeroOrMore(parse.String("A")),
			expectedMatch: nil,
			expectedErr:   nil,
		},
		{
			name:          "ZeroOrMore: something to get",
			input:         "AA",
			parser:        parse.ZeroOrMore(parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedErr:   nil,
		},
		{
			name:          "OneOrMore: nothing to get",
			input:         "BB",
			parser:        parse.OneOrMore(parse.String("A")),
			expectedMatch: nil,
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "OneOrMore: something to get",
			input:         "AA",
			parser:        parse.OneOrMore(parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
