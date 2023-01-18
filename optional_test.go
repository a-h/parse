package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestOptional(t *testing.T) {
	tests := []ParserTest[parse.Match[string]]{
		{
			name:   "Optional: it's not there, but that's OK",
			input:  "ABCDEF",
			parser: parse.Optional(parse.String("1")),
			expectedMatch: parse.Match[string]{
				Value: "",
				OK:    false,
			},
			expectedErr: nil,
		},
		{
			name:   "Optional: it's there, so return the value",
			input:  "ABCDEF",
			parser: parse.Optional(parse.String("A")),
			expectedMatch: parse.Match[string]{
				Value: "A",
				OK:    true,
			},
			expectedErr: nil,
		},
	}
	RunParserTests(t, tests)
}
