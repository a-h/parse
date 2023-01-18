package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestEOF(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:        "no match",
			input:       "A",
			parser:      parse.EOF[string](),
			expectedErr: parse.ErrNotMatched,
		},
		{
			name:        "match",
			input:       "",
			parser:      parse.EOF[string](),
			expectedErr: nil,
		},
	}
	RunParserTests(t, tests)
}
