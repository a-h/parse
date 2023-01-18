package parse_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/a-h/parse"
)

func BenchmarkCSV(b *testing.B) {
	b.ReportAllocs()
	csvData := `a,b,c
"d",e,f
"a string",123,456`
	for i := 0; i < b.N; i++ {
		input := parse.NewInput(csvData)
		_, err := row.Parse(input)
		if err != nil {
			b.Fatalf("failed to parse: %v", err)
		}
	}
}

var colDelimiter = parse.Rune(',')
var doubleQuote = parse.Rune('"')
var escape = parse.Rune('\\')
var escapedQuote = parse.String(`\"`)
var stringUntilEscapedCharacterOrDoubleQuote = parse.StringUntil(parse.Any(escape, doubleQuote))

type QuotedStringParser struct{}

func (p QuotedStringParser) Parse(in parse.Input) (match string, err error) {
	start := in.Position()
	// Start with a quote.
	_, err = doubleQuote.Parse(in)
	if err != nil {
		// No match, so rewind.
		in.Seek(start.Index)
		return
	}
	// Grab the contents.
	var sb strings.Builder
	for {
		// Try for an escaped quote.
		_, err = escapedQuote.Parse(in)
		if err != nil && !errors.Is(err, parse.ErrNotMatched) {
			return
		}
		if !errors.Is(err, parse.ErrNotMatched) {
			sb.WriteRune('"')
			continue
		}
		// Or a terminating quote.
		_, err = doubleQuote.Parse(in)
		if err != nil && !errors.Is(err, parse.ErrNotMatched) {
			return
		}
		if !errors.Is(err, parse.ErrNotMatched) {
			break
		}
		// Grab the runes.
		match, err = stringUntilEscapedCharacterOrDoubleQuote.Parse(in)
		if err != nil && !errors.Is(err, parse.ErrNotMatched) {
			return
		}
		if !errors.Is(err, parse.ErrNotMatched) {
			sb.WriteString(match)
			continue
		}
		// If we haven't gotten a match, we must have reached the end of the file.
		// Without closing the string.
		err = parse.Error("unterminated quoted string", start)
	}
	match = sb.String()
	return
}

var quotedString = parse.Parser[string](QuotedStringParser{})
var rowDelimiter = parse.NewLine
var unquotedString = parse.StringUntil(parse.Any(colDelimiter, rowDelimiter, parse.EOF[string]()))
var stringValueParser = parse.Func(func(in parse.Input) (match string, err error) {
	match, err = parse.Any(quotedString, unquotedString).Parse(in)
	// Chomp the col delimiter, but we could also be at the end of the row, or file.
	parse.Any(colDelimiter, rowDelimiter, parse.EOF[string]()).Parse(in)
	return
})

var row parse.Parser[[]string] = parse.Func(func(in parse.Input) (match []string, err error) {
	match, err = parse.UntilEOF(stringValueParser, rowDelimiter).Parse(in)
	// Chomp the row terminator.
	rowDelimiter.Parse(in)
	return
})

var CSV parse.Parser[[][]string] = parse.UntilEOF(row, parse.Any(colDelimiter, parse.NewLine))
