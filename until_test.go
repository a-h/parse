package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestUntil(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			name:          "Until: success",
			input:         "ABCDEF",
			parser:        parse.Until(parse.AnyRune, parse.String("D")),
			expectedMatch: []string{"A", "B", "C"},
			expectedOK:    true,
		},
		{
			name:       "Until: fail, reached EOF before delimiter was found",
			input:      "ABCDEF",
			parser:     parse.Until(parse.AnyRune, parse.String("G")),
			expectedOK: false,
		},
		{
			name:          "UntilEOF: stop at the delimiter if it's there",
			input:         "ABCDEF",
			parser:        parse.UntilEOF(parse.AnyRune, parse.String("F")),
			expectedMatch: []string{"A", "B", "C", "D", "E"},
			expectedOK:    true,
		},
		{
			name:          "UntilEOF: allow EOF",
			input:         "ABCDEF",
			parser:        parse.UntilEOF(parse.AnyRune, parse.String("G")),
			expectedMatch: []string{"A", "B", "C", "D", "E", "F"},
			expectedOK:    true,
		},
		{
			name:        "Until: return an error on primary failure",
			input:       "ABCDEF",
			parser:      parse.Until(parse.Parser[string](expectErrorParser{}), parse.AnyRune),
			expectedErr: errTestParseError,
		},
		{
			name:        "Until: return an error on delimiter failure",
			input:       "ABCDEF",
			parser:      parse.Until(parse.AnyRune, parse.Parser[string](expectErrorParser{})),
			expectedErr: errTestParseError,
		},
	}
	RunParserTests(t, tests)
}
