package parse

import (
	"testing"
)

func TestZeroToNine(t *testing.T) {
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
			parser:        ZeroToNine,
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "match",
			input:         "0",
			parser:        ZeroToNine,
			expectedMatch: "0",
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
