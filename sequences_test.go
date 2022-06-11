package parse_test

import (
	"testing"

	"github.com/a-h/parse"
)

func TestSequence2(t *testing.T) {
	tests := []ParserTest[parse.Tuple2[string, string]]{
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
			expectedMatch: parse.Tuple2[string, string]{
				A: "ABC",
				B: "DEF",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence3(t *testing.T) {
	tests := []ParserTest[parse.Tuple3[string, string, string]]{
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
			expectedMatch: parse.Tuple3[string, string, string]{
				A: "AB",
				B: "CD",
				C: "EF",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence4(t *testing.T) {
	tests := []ParserTest[parse.Tuple4[string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf4(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf4(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D")),
			expectedMatch: parse.Tuple4[string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence5(t *testing.T) {
	tests := []ParserTest[parse.Tuple5[string, string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf5(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4"), parse.String("5")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf5(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D"), parse.String("E")),
			expectedMatch: parse.Tuple5[string, string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
				E: "E",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence6(t *testing.T) {
	tests := []ParserTest[parse.Tuple6[string, string, string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf6(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4"), parse.String("5"), parse.String("6")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf6(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D"), parse.String("E"), parse.String("F")),
			expectedMatch: parse.Tuple6[string, string, string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
				E: "E",
				F: "F",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence7(t *testing.T) {
	tests := []ParserTest[parse.Tuple7[string, string, string, string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf7(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4"), parse.String("5"), parse.String("6"), parse.String("7")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf7(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D"), parse.String("E"), parse.String("F"), parse.String("G")),
			expectedMatch: parse.Tuple7[string, string, string, string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
				E: "E",
				F: "F",
				G: "G",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence8(t *testing.T) {
	tests := []ParserTest[parse.Tuple8[string, string, string, string, string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf8(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4"), parse.String("5"), parse.String("6"), parse.String("7"), parse.String("8")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf8(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D"), parse.String("E"), parse.String("F"), parse.String("G"), parse.String("H")),
			expectedMatch: parse.Tuple8[string, string, string, string, string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
				E: "E",
				F: "F",
				G: "G",
				H: "H",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}

func TestSequence9(t *testing.T) {
	tests := []ParserTest[parse.Tuple9[string, string, string, string, string, string, string, string, string]]{
		{
			name:       "no match",
			input:      "ABCDEFGHIJ",
			parser:     parse.SequenceOf9(parse.String("1"), parse.String("2"), parse.String("3"), parse.String("4"), parse.String("5"), parse.String("6"), parse.String("7"), parse.String("8"), parse.String("9")),
			expectedOK: false,
		},
		{
			name:   "match",
			input:  "ABCDEFGHIJ",
			parser: parse.SequenceOf9(parse.String("A"), parse.String("B"), parse.String("C"), parse.String("D"), parse.String("E"), parse.String("F"), parse.String("G"), parse.String("H"), parse.String("I")),
			expectedMatch: parse.Tuple9[string, string, string, string, string, string, string, string, string]{
				A: "A",
				B: "B",
				C: "C",
				D: "D",
				E: "E",
				F: "F",
				G: "G",
				H: "H",
				I: "I",
			},
			expectedOK: true,
		},
	}
	RunParserTests(t, tests)
}
