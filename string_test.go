package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestString(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.String("123"),
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "matches",
			input:         "ABCDEF",
			parser:        parse.String("ABC"),
			expectedMatch: "ABC",
			expectedErr:   nil,
		},
		{
			name:          "matches insensitive",
			input:         "ABCDEF",
			parser:        parse.StringInsensitive("abc"),
			expectedMatch: "ABC",
			expectedErr:   nil,
		},
		{
			name:          "matches insensitive (inverse)",
			input:         "abCDEF",
			parser:        parse.StringInsensitive("ABC"),
			expectedMatch: "abC",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
