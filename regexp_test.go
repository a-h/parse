package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestRegexpParser(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match if the regexp doesn't match the start of the string",
			input:         "ABCDEF",
			parser:        parse.MustRegexp("BCD"),
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "matches the start of the string",
			input:         "ABCDEF",
			parser:        parse.MustRegexp("A"),
			expectedMatch: "A",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
