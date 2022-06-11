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
			name:          "times: matches",
			input:         "AAAA",
			parser:        Times(3, String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:          "repeat: must be at least 1, and take up to 5",
			input:         "AAAA",
			parser:        Repeat(1, 5, String("A")),
			expectedMatch: []string{"A", "A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:       "repeat: min of 4, max of 5 - no match",
			input:      "AAA",
			parser:     Repeat(4, 5, String("A")),
			expectedOK: false,
		},
		{
			name:          "repeat: min of 0, max of 2 - matches",
			input:         "AAA",
			parser:        Repeat(0, 2, String("A")),
			expectedMatch: []string{"A", "A"},
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
