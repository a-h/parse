package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestTimes(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			name:       "Times: no match",
			input:      "ABCDEF",
			parser:     parse.Times(2, parse.String("A")),
			expectedOK: false,
		},
		{
			name:          "Times: matches",
			input:         "AAAA",
			parser:        parse.Times(3, parse.String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:          "Repeat: must be at least 1, and take up to 5",
			input:         "AAAA",
			parser:        parse.Repeat(1, 5, parse.String("A")),
			expectedMatch: []string{"A", "A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:       "Repeat: min of 4, max of 5 - no match",
			input:      "AAA",
			parser:     parse.Repeat(4, 5, parse.String("A")),
			expectedOK: false,
		},
		{
			name:          "Repeat: min of 0, max of 2 - matches",
			input:         "AAA",
			parser:        parse.Repeat(0, 2, parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "AtMost: success",
			input:         "AAA",
			parser:        parse.AtMost(2, parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "AtLeast: success",
			input:         "AAA",
			parser:        parse.AtLeast(2, parse.String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:          "ZeroOrMore: nothing to get",
			input:         "BB",
			parser:        parse.ZeroOrMore(parse.String("A")),
			expectedMatch: nil,
			expectedOK:    true,
		},
		{
			name:          "ZeroOrMore: something to get",
			input:         "AA",
			parser:        parse.ZeroOrMore(parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "OneOrMore: nothing to get",
			input:         "BB",
			parser:        parse.OneOrMore(parse.String("A")),
			expectedMatch: nil,
			expectedOK:    false,
		},
		{
			name:          "OneOrMore: something to get",
			input:         "AA",
			parser:        parse.OneOrMore(parse.String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}
