package parse_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/a-h/parse"
)

func TestRow(t *testing.T) {
	csvData := `a,b,c
`
	input := parse.NewInput(csvData)
	row, ok, err := row.Parse(input)
	if len(row) != 3 || !ok {
		t.Errorf("failed to parse: %v, %v", row, err)
	}
}

func TestCSV(t *testing.T) {
	csvData := `a,b,c
"d",e,f
"a string",123,456`
	input := parse.NewInput(csvData)
	rows, ok, err := CSV.Parse(input)
	if err != nil || !ok {
		t.Fatalf("failed to parse: %v", err)
	}
	if len(rows) != 3 {
		t.Errorf("expected 3 rows, got %d", len(rows))
		j, _ := json.Marshal(rows)
		t.Error(string(j))
	}
}

func BenchmarkCSV(b *testing.B) {
	b.ReportAllocs()
	csvData := `a,b,c
"d",e,f
"a string",123,456`
	for i := 0; i < b.N; i++ {
		input := parse.NewInput(csvData)
		_, ok, err := CSV.Parse(input)
		if err != nil || !ok {
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

func (p QuotedStringParser) Parse(in *parse.Input) (match string, ok bool, err error) {
	start := in.Position()
	// Start with a quote.
	_, ok, err = doubleQuote.Parse(in)
	if err != nil || !ok {
		// No match, so rewind.
		in.Seek(start.Index)
		return
	}
	// Grab the contents.
	var sb strings.Builder
	for {
		// Try for an escaped quote.
		_, ok, err = escapedQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
			sb.WriteRune('"')
			continue
		}
		// Or a terminating quote.
		_, ok, err = doubleQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
			break
		}
		// Grab the runes.
		match, ok, err = stringUntilEscapedCharacterOrDoubleQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
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
var unquotedString = parse.StringUntilEOF(parse.Any(colDelimiter, rowDelimiter))
var stringValueParser = parse.Func(func(in *parse.Input) (match string, ok bool, err error) {
	match, ok, err = parse.Any(quotedString, unquotedString).Parse(in)
	// Chomp the col delimiter.
	colDelimiter.Parse(in)
	return
})

var row parse.Parser[[]string] = parse.Func(func(in *parse.Input) (match []string, ok bool, err error) {
	match, ok, err = parse.UntilEOF(stringValueParser, rowDelimiter).Parse(in)
	rowDelimiter.Parse(in)
	return
})

var CSV parse.Parser[[][]string] = parse.Until(row, parse.EOF[string]())
