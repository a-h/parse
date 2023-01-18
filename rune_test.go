package parse_test

import (
	"testing"
	"unicode"

	"github.com/a-h/parse"
)

func TestRuneWhere(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "RuneWhere: no match",
			input:         "ABCDEF",
			parser:        parse.RuneWhere(func(r rune) bool { return r == 'a' }),
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:  "RuneWhere: match",
			input: "ABCDEF",
			parser: parse.RuneWhere(func(r rune) bool {
				return unicode.IsUpper(r)
			}),
			expectedMatch: "A",
			expectedErr:   nil,
		},
		{
			name:          "AnyRune: match",
			input:         "ABCDEF",
			parser:        parse.AnyRune,
			expectedMatch: "A",
			expectedErr:   nil,
		},
		{
			name:          "RuneIn: no match",
			input:         "ABCDEF",
			parser:        parse.RuneIn("123"),
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "RuneIn: match",
			input:         "ABCDEF",
			parser:        parse.RuneIn("CBA"),
			expectedMatch: "A",
			expectedErr:   nil,
		},
		{
			name:          "RuneNotIn: no match",
			input:         "ABCDEF",
			parser:        parse.RuneNotIn("ABC"),
			expectedMatch: "",
			expectedErr:   parse.ErrNotMatched,
		},
		{
			name:          "RuneNotIn: match",
			input:         "ABCDEF",
			parser:        parse.RuneNotIn("123"),
			expectedMatch: "A",
			expectedErr:   nil,
		},
		{
			name:          "RuneInRanges: match",
			input:         "     ",
			parser:        parse.RuneInRanges(unicode.White_Space),
			expectedMatch: " ",
			expectedErr:   nil,
		},
		{
			name:        "RuneInRanges: no match",
			input:       "     ",
			parser:      parse.RuneInRanges(unicode.Han),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:          "Letter: match",
			input:         "a",
			parser:        parse.Letter,
			expectedMatch: "a",
			expectedErr:   nil,
		},
	}
	RunParserTests(t, tests)
}
