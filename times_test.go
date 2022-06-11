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
			name:       "Times: no match",
			input:      "ABCDEF",
			parser:     Times(2, String("A")),
			expectedOK: false,
		},
		{
			name:          "Times: matches",
			input:         "AAAA",
			parser:        Times(3, String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:          "Repeat: must be at least 1, and take up to 5",
			input:         "AAAA",
			parser:        Repeat(1, 5, String("A")),
			expectedMatch: []string{"A", "A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:       "Repeat: min of 4, max of 5 - no match",
			input:      "AAA",
			parser:     Repeat(4, 5, String("A")),
			expectedOK: false,
		},
		{
			name:          "Repeat: min of 0, max of 2 - matches",
			input:         "AAA",
			parser:        Repeat(0, 2, String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "AtMost: success",
			input:         "AAA",
			parser:        AtMost(2, String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "AtLeast: success",
			input:         "AAA",
			parser:        AtLeast(2, String("A")),
			expectedMatch: []string{"A", "A", "A"},
			expectedOK:    true,
		},
		{
			name:          "ZeroOrMore: nothing to get",
			input:         "BB",
			parser:        ZeroOrMore(String("A")),
			expectedMatch: nil,
			expectedOK:    true,
		},
		{
			name:          "ZeroOrMore: something to get",
			input:         "AA",
			parser:        ZeroOrMore(String("A")),
			expectedMatch: []string{"A", "A"},
			expectedOK:    true,
		},
		{
			name:          "OneOrMore: nothing to get",
			input:         "BB",
			parser:        OneOrMore(String("A")),
			expectedMatch: nil,
			expectedOK:    false,
		},
		{
			name:          "OneOrMore: something to get",
			input:         "AA",
			parser:        OneOrMore(String("A")),
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
