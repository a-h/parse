package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestEOF(t *testing.T) {
	tests := []ParserTest[string]{
		{
			name:       "no match",
			input:      "A",
			parser:     parse.EOF[string](),
			expectedOK: false,
		},
		{
			name:       "match",
			input:      "",
			parser:     parse.EOF[string](),
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
