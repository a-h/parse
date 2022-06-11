package parse

import (
	"testing"
)

func TestRegexpParser(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[string]
		expectedMatch string
		expectedOK    bool
	}{
		{
			name:          "no match if the regexp doesn't match the start of the string",
			input:         "ABCDEF",
			parser:        MustRegexp("BCD"),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "matches the start of the string",
			input:         "ABCDEF",
			parser:        MustRegexp("A"),
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
