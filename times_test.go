package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTimes(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[[]string]
		expectedMatch []string
		expectedOK    bool
	}{
		{
			name:       "no match",
			input:      "ABCDEF",
			parser:     Times(2, String("A")),
			expectedOK: false,
		},
		{
			name:          "matches",
			input:         "AAAA",
			parser:        Times(3, String("A")),
			expectedMatch: []string{"A", "A", "A"},
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
		if !test.expectedOK {
			continue
		}
		if diff := cmp.Diff(test.expectedMatch, match); diff != "" {
			t.Error(diff)
		}
	}
}
