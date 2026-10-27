package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestOr(t *testing.T) {
	tests := []ParserTest[parse.Tuple2[parse.Result[string], parse.Result[string]]]{
		{
			name:          "no match",
			input:         "C",
			parser:        parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Result[string], parse.Result[string]]{},
			expectedOK:    false,
		},
		{
			name:   "match",
			input:  "A",
			parser: parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Result[string], parse.Result[string]]{
				A: parse.Result[string]{
					Value: "A",
					OK:    true,
				},
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
