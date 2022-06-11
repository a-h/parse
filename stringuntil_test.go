package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringUntil(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[string]
		expectedMatch string
		expectedOK    bool
	}{
		{
			name:          "StringUntil: success",
			input:         "ABCDEF",
			parser:        StringUntil(String("D")),
			expectedMatch: "ABC",
			expectedOK:    true,
		},
		{
			name:       "StringUntil: fail, reached EOF before delimiter was found",
			input:      "ABCDEF",
			parser:     StringUntil(String("G")),
			expectedOK: false,
		},
		{
			name:          "StringUntilEOF: stop at the delimiter if it's there",
			input:         "ABCDEF",
			parser:        StringUntilEOF(String("F")),
			expectedMatch: "ABCDE",
			expectedOK:    true,
		},
		{
			name:          "StringUntilEOF: allow EOF",
			input:         "ABCDEF",
			parser:        StringUntilEOF(String("G")),
			expectedMatch: "ABCDEF",
			expectedOK:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := NewInput(test.input)
			match, ok, err := test.parser.Parse(in)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}
			if ok != test.expectedOK {
				t.Errorf("expected ok=%v, got=%v", test.expectedOK, ok)
			}
			if !test.expectedOK {
				return
			}
			if diff := cmp.Diff(test.expectedMatch, match); diff != "" {
				t.Error(diff)
			}
		})
	}
}
