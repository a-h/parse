package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestEOF(t *testing.T) {
	tests := []ParserTest[any]{
		{
			name:       "no match",
			input:      "A",
			parser:     parse.EOF,
			expectedOK: false,
		},
		{
			name:       "match",
			input:      "",
			parser:     parse.EOF,
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
