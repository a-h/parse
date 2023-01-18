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
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:   "match A",
			input:  "A",
			parser: parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Match[string], parse.Match[string]]{
				A: parse.Match[string]{
					Value: "A",
					OK:    true,
				},
			},
			expectedErr: nil,
		},
		{
			name:   "match B",
			input:  "B",
			parser: parse.Or(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: parse.Tuple2[parse.Match[string], parse.Match[string]]{
				B: parse.Match[string]{
					Value: "B",
					OK:    true,
				},
			},
			expectedErr: nil,
		},
	}
	RunParserTests(t, tests)
}
