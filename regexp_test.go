package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
