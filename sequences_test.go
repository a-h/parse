package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestSequence2(t *testing.T) {
	tests := []ParserTest[parse.SequenceOf2Result[string, string]]{
		{
			name:       "no match",
			input:      "ABCDEF",
			parser:     parse.SequenceOf2(parse.String("123"), parse.String("ABC")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEF",
			parser: parse.SequenceOf2(parse.String("ABC"), parse.String("DEF")),
			expectedMatch: parse.SequenceOf2Result[string, string]{
				A: "ABC",
				B: "DEF",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence3(t *testing.T) {
	tests := []ParserTest[parse.SequenceOf3Result[string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEF",
			parser:     parse.SequenceOf3(parse.String("12"), parse.String("34"), parse.String("56")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEF",
			parser: parse.SequenceOf3(parse.String("AB"), parse.String("CD"), parse.String("EF")),
			expectedMatch: parse.SequenceOf3Result[string, string, string]{
				A: "AB",
				B: "CD",
				C: "EF",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
