package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestOr(t *testing.T) {
	tests := []ParserTest[parse.OrResult[string, string]]{
		{
			name:          "no match",
			input:         "C",
			parser:        parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.OrResult[string, string]{},
			expectedOK:    false,
		},
		{
			name:   "match",
			input:  "A",
			parser: parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.OrResult[string, string]{
				A: parse.OptionalResult[string]{
					Value: "A",
					OK:    true,
				},
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
