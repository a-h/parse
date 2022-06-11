package parse_test

import (
	"testing"

	"github.com/a-h/parse"
	"github.com/google/go-cmp/cmp"
)

type ParserTest[T any] struct {
	name          string
	input         string
	parser        parse.Parser[T]
	expectedMatch T
	expectedOK    bool
}

func RunParserTests[T any](t *testing.T, tests []ParserTest[T]) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := parse.NewInput(test.input)
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
