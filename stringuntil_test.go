package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestStringUntil(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "StringUntil: success",
			input:         "ABCDEF",
			parser:        parse.StringUntil(parse.String("D")),
			expectedMatch: "ABC",
			expectedErr:   nil,
		},
		{
			name:        "StringUntil: fail, reached EOF before delimiter was found",
			input:       "ABCDEF",
			parser:      parse.StringUntil(parse.String("G")),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "StringUntilEOF: stop at the delimiter if it's there",
			input:         "ABCDEF",
			parser:        parse.StringUntilEOF(parse.String("F")),
			expectedMatch: "ABCDE",
			expectedErr:   nil,
		},
		{
			name:          "StringUntilEOF: allow EOF",
			input:         "ABCDEF",
			parser:        parse.StringUntilEOF(parse.String("G")),
			expectedMatch: "ABCDEF",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
