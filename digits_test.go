package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestZeroToNine(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.ZeroToNine,
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "match",
			input:         "0",
			parser:        parse.ZeroToNine,
			expectedMatch: "0",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
