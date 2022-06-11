package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
