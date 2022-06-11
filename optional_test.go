package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOptional(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[OptionalResult[string]]
		expectedMatch OptionalResult[string]
		expectedOK    bool
	}{
		{
			name:   "Optional: it's not there, but that's OK",
			input:  "ABCDEF",
			parser: Optional(String("1")),
			expectedMatch: OptionalResult[string]{
				Value: "",
				OK:    false,
			},
			expectedOK: true,
		},
		{
			name:   "Optional: it's there, so return the value",
			input:  "ABCDEF",
			parser: Optional(String("A")),
			expectedMatch: OptionalResult[string]{
				Value: "A",
				OK:    true,
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
