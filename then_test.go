package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestThen(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[SequenceOf2Result[string, string]]
		expectedMatch SequenceOf2Result[string, string]
		expectedOK    bool
	}{
		{
			name:       "no match",
			input:      "ABCDEF",
			parser:     Then(String("ABC"), String("456")),
			expectedOK: false,
		},
		{
			name:   "matches",
			input:  "ABCDEF",
			parser: Then(String("ABC"), String("DEF")),
			expectedMatch: SequenceOf2Result[string, string]{
				A: "ABC",
				B: "DEF",
			},
			expectedOK: true,
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
