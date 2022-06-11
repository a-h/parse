package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrParser(t *testing.T) {
	in := NewInput("A")
	matchA := String("A")
	matchB := String("B")
	matchEither := Or(matchA, matchB)
	item, ok, err := matchEither.Parse(in)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}
	if !ok {
		t.Errorf("expected match, but didn't")
	}
	if !item.A.OK {
		t.Errorf("expected A to match, but didn't")
	}
	if item.B.OK {
		t.Errorf("expected B not to match, but it did")
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[OrResult[string, string]]
		expectedMatch OrResult[string, string]
		expectedOK    bool
	}{
		{
			name:          "no match",
			input:         "C",
			parser:        Or(Rune('A'), Rune('B')),
			expectedMatch: OrResult[string, string]{},
			expectedOK:    false,
		},
		{
			name:   "match",
			input:  "A",
			parser: Or(Rune('A'), Rune('B')),
			expectedMatch: OrResult[string, string]{
				A: Option[string]{
					Value: "A",
					OK:    true,
				},
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
		if !test.expectedOK {
			continue
		}
		if diff := cmp.Diff(test.expectedMatch, match); diff != "" {
			t.Error(diff)
		}
	}
}
