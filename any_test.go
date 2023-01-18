package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestAny(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:        "no match",
			input:       "C",
			parser:      parse.Any(parse.Rune('A'), parse.Rune('B')),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "match",
			input:         "B",
			parser:        parse.Any(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: "B",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
