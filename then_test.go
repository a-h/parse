package parse

import (
	"testing"
)

func TestThen(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		parser        Parser[ThenResult[string, string]]
		expectedMatch ThenResult[string, string]
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
			expectedMatch: ThenResult[string, string]{
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
