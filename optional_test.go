package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestOptional(t *testing.T) {
	tests := []ParserTest[parse.Result[string]]{
		{
			name:   "Optional: it's not there, but that's OK",
			input:  "ABCDEF",
			parser: parse.Optional(parse.String("1")),
			expectedMatch: parse.Result[string]{
				Value: "",
				OK:    false,
			},
			expectedOK: true,
		},
		{
			name:   "Optional: it's there, so return the value",
			input:  "ABCDEF",
			parser: parse.Optional(parse.String("A")),
			expectedMatch: parse.Result[string]{
				Value: "A",
				OK:    true,
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
