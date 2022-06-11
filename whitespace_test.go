package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWhitespace(t *testing.T) {
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
			parser:        Whitespace,
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name: "match",
			input: " 	ABC",
			parser: Whitespace,
			expectedMatch: " 	",
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

func TestOptionalWhitespace(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[OptionalResult[string]]
		expectedMatch OptionalResult[string]
		expectedOK    bool
	}{
		{
			name:          "no match",
			input:         "ABCDEF",
			parser:        OptionalWhitespace,
			expectedMatch: OptionalResult[string]{},
			expectedOK:    true,
		},
		{
			name: "match",
			input: " 	ABC",
			parser: OptionalWhitespace,
			expectedMatch: OptionalResult[string]{
				Value: " 	",
				OK: true,
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
