package parse

import (
	"testing"
	"unicode"
)

func TestRuneWhere(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[string]
		expectedMatch string
		expectedOK    bool
	}{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        RuneWhere(func(r rune) bool { return r == 'a' }),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:  "match",
			input: "ABCDEF",
			parser: RuneWhere(func(r rune) bool {
				return unicode.IsUpper(r)
			}),
			expectedMatch: "A",
			expectedOK:    true,
		},
	}

	for _, test := range tests {
		in := NewInput(test.input)
		match, ok, err := test.parser.Parse(in)
		if err != nil {
			t.Fatalf("failed to parse: %v", err)
		}
		if ok != test.expectedOK {
			t.Errorf("expected ok=%v, got=%v", test.expectedOK, ok)
		}
		if test.expectedOK && match != test.expectedMatch {
			t.Errorf("expected match=%q, got=%q", test.expectedMatch, match)
		}
	}
}
