package parse_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/a-h/parse"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type matchResult struct {
	OK    bool
	Err   error
	Index int
}

func matchInput(input string, m parse.Matcher) (r matchResult) {
	in := parse.NewInput([]byte(input))
	r.OK, r.Err = m.Match(in)
	r.Index = in.Index()
	return
}

func createParseFunc[T any](p parse.Parser[T]) func(in *parse.Input) (ok bool, err error) {
	return func(in *parse.Input) (ok bool, err error) {
		_, ok, err = p.Parse(in)
		return
	}
}

func parseInput(input string, parseFunc func(in *parse.Input) (ok bool, err error)) (r matchResult) {
	in := parse.NewInput([]byte(input))
	r.OK, r.Err = parseFunc(in)
	r.Index = in.Index()
	return
}

func TestMatchAgreesWithParse(t *testing.T) {
	errConvert := errors.New("convert failed")
	failing := parse.Parser[string](expectErrorParser{})
	tests := []struct {
		name    string
		input   string
		matcher parse.Matcher
		parse   func(in *parse.Input) (ok bool, err error)
	}{
		{name: "String matches", input: "abc", matcher: parse.String("ab"), parse: createParseFunc(parse.String("ab"))},
		{name: "String does not match", input: "abc", matcher: parse.String("b"), parse: createParseFunc(parse.String("b"))},
		{name: "StringInsensitive matches", input: "ABC", matcher: parse.StringInsensitive("ab"), parse: createParseFunc(parse.StringInsensitive("ab"))},
		{name: "ASCII Rune matches", input: "<a", matcher: parse.Rune('<'), parse: createParseFunc(parse.Rune('<'))},
		{name: "multi-byte Rune matches", input: "éa", matcher: parse.Rune('é'), parse: createParseFunc(parse.Rune('é'))},
		{name: "RuneIn matches", input: "7", matcher: parse.ZeroToNine, parse: createParseFunc(parse.ZeroToNine)},
		{name: "Regexp matches", input: "abc123", matcher: parse.MustRegexp("[a-z]+"), parse: createParseFunc(parse.MustRegexp("[a-z]+"))},
		{name: "EOF matches the end of the input", input: "", matcher: parse.EOF[string](), parse: createParseFunc(parse.EOF[string]())},
		{name: "EOF does not match other input", input: "a", matcher: parse.EOF[string](), parse: createParseFunc(parse.EOF[string]())},
		{name: "All matches", input: "abc", matcher: parse.All(parse.Rune('a'), parse.Rune('b')), parse: createParseFunc(parse.All(parse.Rune('a'), parse.Rune('b')))},
		{name: "All rolls back a partial match", input: "ac", matcher: parse.All(parse.Rune('a'), parse.Rune('b')), parse: createParseFunc(parse.All(parse.Rune('a'), parse.Rune('b')))},
		{name: "All returns errors", input: "a", matcher: parse.All(parse.Rune('a'), failing), parse: createParseFunc(parse.All(parse.Rune('a'), failing))},
		{name: "Any matches the second parser", input: "b", matcher: parse.Any(parse.Rune('a'), parse.Rune('b')), parse: createParseFunc(parse.Any(parse.Rune('a'), parse.Rune('b')))},
		{name: "Or matches the second parser", input: "1", matcher: parse.Or(parse.Rune('a'), parse.ZeroToNine), parse: createParseFunc(parse.Or(parse.Rune('a'), parse.ZeroToNine))},
		{name: "Optional matches nothing", input: "b", matcher: parse.Optional(parse.Rune('a')), parse: createParseFunc(parse.Optional(parse.Rune('a')))},
		{name: "Times rolls back too few matches", input: "aab", matcher: parse.Times(3, parse.Rune('a')), parse: createParseFunc(parse.Times(3, parse.Rune('a')))},
		{name: "Repeat stops at the maximum", input: "aaaa", matcher: parse.Repeat(1, 2, parse.Rune('a')), parse: createParseFunc(parse.Repeat(1, 2, parse.Rune('a')))},
		{name: "ZeroOrMore matches nothing", input: "b", matcher: parse.ZeroOrMore(parse.Rune('a')), parse: createParseFunc(parse.ZeroOrMore(parse.Rune('a')))},
		{name: "Times returns errors", input: "a", matcher: parse.Times(2, failing), parse: createParseFunc(parse.Times(2, failing))},
		{name: "StringFrom matches", input: "ab;", matcher: parse.StringFrom(parse.Letter, parse.Letter), parse: createParseFunc(parse.StringFrom(parse.Letter, parse.Letter))},
		{name: "StringUntil matches", input: "ab;", matcher: parse.StringUntil(parse.Rune(';')), parse: createParseFunc(parse.StringUntil(parse.Rune(';')))},
		{name: "StringUntil does not match without a delimiter", input: "ab", matcher: parse.StringUntil(parse.Rune(';')), parse: createParseFunc(parse.StringUntil(parse.Rune(';')))},
		{name: "StringUntilEOF matches to the end", input: "ab", matcher: parse.StringUntilEOF(parse.Rune(';')), parse: createParseFunc(parse.StringUntilEOF(parse.Rune(';')))},
		{name: "StringUntil returns delimiter errors", input: "ab", matcher: parse.StringUntil(failing), parse: createParseFunc(parse.StringUntil(failing))},
		{name: "Until matches", input: "ab;", matcher: parse.Until(parse.Letter, parse.Rune(';')), parse: createParseFunc(parse.Until(parse.Letter, parse.Rune(';')))},
		{name: "Until rolls back when an item does not match", input: "a1;", matcher: parse.Until(parse.Letter, parse.Rune(';')), parse: createParseFunc(parse.Until(parse.Letter, parse.Rune(';')))},
		{name: "UntilEOF matches to the end", input: "ab", matcher: parse.UntilEOF(parse.Letter, parse.Rune(';')), parse: createParseFunc(parse.UntilEOF(parse.Letter, parse.Rune(';')))},
		{name: "Until returns item errors", input: "ab", matcher: parse.Until(failing, parse.Rune(';')), parse: createParseFunc(parse.Until(failing, parse.Rune(';')))},
		{name: "SequenceOf2 does not match", input: "a1", matcher: parse.SequenceOf2(parse.Letter, parse.Letter), parse: createParseFunc(parse.SequenceOf2(parse.Letter, parse.Letter))},
		{name: "SequenceOf3 rolls back a partial match", input: "ab1", matcher: parse.SequenceOf3(parse.Letter, parse.Letter, parse.Letter), parse: createParseFunc(parse.SequenceOf3(parse.Letter, parse.Letter, parse.Letter))},
		{name: "SequenceOf9 matches", input: "abcdefghi", matcher: parse.SequenceOf9(parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter), parse: createParseFunc(parse.SequenceOf9(parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter, parse.Letter))},
		{name: "Then matches", input: "ab", matcher: parse.Then(parse.Letter, parse.Letter), parse: createParseFunc(parse.Then(parse.Letter, parse.Letter))},
		{name: "Whitespace matches", input: " \t\nx", matcher: parse.Whitespace, parse: createParseFunc(parse.Whitespace)},
		{name: "OptionalWhitespace matches nothing", input: "x", matcher: parse.OptionalWhitespace, parse: createParseFunc(parse.OptionalWhitespace)},
		{
			name:    "Convert returns conversion errors",
			input:   "1",
			matcher: parse.Convert(parse.ZeroToNine, func(s string) (int, error) { return 0, errConvert }),
			parse:   createParseFunc(parse.Convert(parse.ZeroToNine, func(s string) (int, error) { return 0, errConvert })),
		},
		{
			name:    "Convert matches",
			input:   "1",
			matcher: parse.Convert(parse.ZeroToNine, strconv.Atoi),
			parse:   createParseFunc(parse.Convert(parse.ZeroToNine, strconv.Atoi)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := parseInput(tt.input, tt.parse)
			actual := matchInput(tt.input, tt.matcher)
			if diff := cmp.Diff(expected, actual, cmpopts.EquateErrors()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
