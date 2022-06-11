package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSequence2(t *testing.T) {
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
			parser:     SequenceOf2(String("123"), String("ABC")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEF",
			parser: SequenceOf2(String("ABC"), String("DEF")),
			expectedMatch: SequenceOf2Result[string, string]{
				A: "ABC",
				B: "DEF",
			},
			expectedOK: true,
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

func TestSequence3(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[SequenceOf3Result[string, string, string]]
		expectedMatch SequenceOf3Result[string, string, string]
		expectedOK    bool
	}{
		{
			name:       "no match",
			input:      "ABCDEF",
			parser:     SequenceOf3(String("12"), String("34"), String("56")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEF",
			parser: SequenceOf3(String("AB"), String("CD"), String("EF")),
			expectedMatch: SequenceOf3Result[string, string, string]{
				A: "AB",
				B: "CD",
				C: "EF",
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
