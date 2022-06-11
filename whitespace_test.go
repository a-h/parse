package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestWhitespace(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.Whitespace,
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name: "match",
			input: " 	ABC",
			parser: parse.Whitespace,
			expectedMatch: " 	",
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestOptionalWhitespace(t *testing.T) {
	tests := []ParserTest[parse.Match[string]]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.OptionalWhitespace,
			expectedMatch: parse.Match[string]{},
			expectedOK:    true,
		},
		{
			name: "match",
			input: " 	ABC",
			parser: parse.OptionalWhitespace,
			expectedMatch: parse.Match[string]{
				Value: " 	",
				OK: true,
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
