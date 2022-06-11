package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestStringFrom(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        parse.StringFrom(parse.String("ABC"), parse.String("123")),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "matches",
			input:         "ABCDEF",
			parser:        parse.StringFrom(parse.MustRegexp("."), parse.MustRegexp("BC"), parse.String("DEF")),
			expectedMatch: "ABCDEF",
			expectedOK:    true,
		},
	}
	RunParserTests(t, tests)
}
