package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestThen(t *testing.T) {
	tests := []ParserTest[parse.Tuple2[string, string]]{
		{
			name:        "no match",
			input:       "ABCDEF",
			parser:      parse.Then(parse.String("ABC"), parse.String("456")),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:   "matches",
			input:  "ABCDEF",
			parser: parse.Then(parse.String("ABC"), parse.String("DEF")),
			expectedMatch: parse.Tuple2[string, string]{
				A: "ABC",
				B: "DEF",
			},
			expectedErr: nil,
		},
	}
	RunParserTests(t, tests)
}
