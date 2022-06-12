package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestAll(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			name:       "no match",
			input:      "AC",
			parser:     parse.All(parse.Rune('A'), parse.Rune('B')),
			expectedOK: false,
		},
		{
			name:          "match",
			input:         "AB",
			parser:        parse.All(parse.Rune('A'), parse.Rune('B')),
			expectedMatch: []string{"A", "B"},
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}
