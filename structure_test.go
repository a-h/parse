package parse_test

import (
	"testing"

	"github.com/a-h/parse"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type ParserTest[T any] struct {
	name          string
	input         string
	parser        parse.Parser[T]
	expectedMatch T
	expectedOK    bool
	expectedErr   error
}

func RunParserTests[T any](t *testing.T, tests []ParserTest[T]) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := parse.NewInput(test.input)
			match, ok, err := test.parser.Parse(in)
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected parser error: %v", err)
			}
			if test.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected err=%v, got nil", test.expectedErr)
				}
				if diff := cmp.Diff(test.expectedErr, err, cmpopts.EquateErrors()); diff != "" {
					t.Errorf("error\n:%s", diff)
				}
				return
			}
			if ok != test.expectedOK {
				t.Errorf("expected ok=%v, got=%v", test.expectedOK, ok)
			}
			if !test.expectedOK {
				if in.Index() != 0 {
					t.Error("input not rolled back")
				}
				return
			}
			if diff := cmp.Diff(test.expectedMatch, match); diff != "" {
				t.Error(diff)
			}
		})
	}
}
