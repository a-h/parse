package parse

import (
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

func TestRuneWhere(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[string]
		expectedMatch string
		expectedOK    bool
	}{
		{
			name:          "RuneWhere: no match",
			input:         "ABCDEF",
			parser:        RuneWhere(func(r rune) bool { return r == 'a' }),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:  "RuneWhere: match",
			input: "ABCDEF",
			parser: RuneWhere(func(r rune) bool {
				return unicode.IsUpper(r)
			}),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "AnyRune: match",
			input:         "ABCDEF",
			parser:        AnyRune,
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneIn: no match",
			input:         "ABCDEF",
			parser:        RuneIn("123"),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "RuneIn: match",
			input:         "ABCDEF",
			parser:        RuneIn("CBA"),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneNotIn: no match",
			input:         "ABCDEF",
			parser:        RuneNotIn("ABC"),
			expectedMatch: "",
			expectedOK:    false,
		},
		{
			name:          "RuneNotIn: match",
			input:         "ABCDEF",
			parser:        RuneNotIn("123"),
			expectedMatch: "A",
			expectedOK:    true,
		},
		{
			name:          "RuneInRanges: match",
			input:         "     ",
			parser:        RuneInRanges(unicode.White_Space),
			expectedMatch: " ",
			expectedOK:    true,
		},
		{
			name:       "RuneInRanges: no match",
			input:      "     ",
			parser:     RuneInRanges(unicode.Han),
			expectedOK: false,
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
