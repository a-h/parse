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
			expectedOK:    false,
		},
		{
			name:          "matches",
			input:         "ABCDEF",
			parser:        parse.String("ABC"),
			expectedMatch: "ABC",
			expectedOK:    true,
		},
		{
			name:          "matches insensitive",
			input:         "ABCDEF",
			parser:        parse.StringInsensitive("abc"),
			expectedMatch: "ABC",
			expectedOK:    true,
		},
		{
			name:          "matches insensitive (inverse)",
			input:         "abCDEF",
			parser:        parse.StringInsensitive("ABC"),
			expectedMatch: "abC",
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}
