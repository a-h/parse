package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestAll(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			name:        "no match",
			input:       "AC",
			parser:      parse.All(parse.Rune('A'), parse.Rune('B')),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "match",
			input:         "AB",
			parser:        parse.All(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: []string{"A", "B"},
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
