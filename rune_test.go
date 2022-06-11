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
			expectedOK:    false,
		},
		{
			name:  "RuneWhere: match",
			input: "ABCDEF",
			parser: parse.RuneWhere(func(r rune) bool {
				return unicode.IsUpper(r)
			}),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "AnyRune: match",
			input:         "ABCDEF",
			parser:        parse.AnyRune,
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneIn: no match",
			input:         "ABCDEF",
			parser:        parse.RuneIn("123"),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "RuneIn: match",
			input:         "ABCDEF",
			parser:        parse.RuneIn("CBA"),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneNotIn: no match",
			input:         "ABCDEF",
			parser:        parse.RuneNotIn("ABC"),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "RuneNotIn: match",
			input:         "ABCDEF",
			parser:        parse.RuneNotIn("123"),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneInRanges: match",
			input:         "     ",
			parser:        parse.RuneInRanges(unicode.White_Space),
			expectedMatch: " ",
			expectedOK:    true,
		},
		{
			name:       "RuneInRanges: no match",
			input:      "     ",
			parser:     parse.RuneInRanges(unicode.Han),
			expectedOK: false,
		},
		{
			name:          "Letter: match",
			input:         "a",
			parser:        parse.Letter,
			expectedMatch: "a",
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}
