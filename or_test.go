package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestOr(t *testing.T) {
	tests := []ParserTest[parse.Tuple2[parse.Match[string], parse.Match[string]]]{
		{
			name:          "no match",
			input:         "C",
			parser:        parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Match[string], parse.Match[string]]{},
			expectedOK:    false,
		},
		{
			name:   "match",
			input:  "A",
			parser: parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Match[string], parse.Match[string]]{
				A: parse.Match[string]{
					Value: "A",
					OK:    true,
				},
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
