package parse_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/parse"
)

func TestRow(t *testing.T) {
	csvData := `a,b,c
`
	input := parse.NewInput([]byte(csvData))
	row, ok, err := row.Parse(input)
	if len(row) != 3 || !ok {
		t.Errorf("failed to parse: %v, %v", row, err)
	}
}

func TestCSV(t *testing.T) {
	csvData := `a,b,c
"d",e,f
"a string",123,456`
	input := parse.NewInput([]byte(csvData))
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
	csvData := []byte(`a,b,c
"d",e,f
"a string",123,456`)
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

func (p QuotedStringParser) Match(in *parse.Input) (ok bool, err error) {
	return parse.Match(p, in)
}

var quotedString = parse.Parser[string](QuotedStringParser{})
var rowDelimiter = parse.NewLine
var unquotedString = parse.StringUntilEOF(parse.Any(colDelimiter, rowDelimiter))
var stringValueParser = parse.Func(func(in *parse.Input) (match string, ok bool, err error) {
	match, ok, err = parse.Any(quotedString, unquotedString).Parse(in)
	// Chomp the col delimiter.
	_, _, _ = colDelimiter.Parse(in)
	return
})

var row parse.Parser[[]string] = parse.Func(func(in *parse.Input) (match []string, ok bool, err error) {
	match, ok, err = parse.UntilEOF(stringValueParser, rowDelimiter).Parse(in)
	_, _, _ = rowDelimiter.Parse(in)
	return
})

var CSV parse.Parser[[][]string] = parse.Until(row, parse.EOF[string]())

func runParserBenchmark[T any](b *testing.B, input string, parser parse.Parser[T]) {
	b.ReportAllocs()
	in := parse.NewInput([]byte(input))
	for b.Loop() {
		in.Seek(0)
		_, ok, err := parser.Parse(in)
		if err != nil || !ok {
			b.Fatalf("failed to parse: ok=%v, err=%v", ok, err)
		}
	}
}

func BenchmarkParseStrings(b *testing.B) {
	benchmarks := []struct {
		name   string
		input  string
		parser parse.Parser[string]
	}{
		{name: "Whitespace", input: "    \t\t    \n    \t\t    \n    x", parser: parse.Whitespace},
		{name: "OptionalWhitespace", input: "    \t\t    \n    \t\t    \n    x", parser: parse.OptionalWhitespace},
		{name: "StringFrom OneOrMore Letter", input: "abcdefghijklmnopqrstuvwxyz ", parser: parse.StringFrom(parse.OneOrMore(parse.Letter))},
		{name: "StringFrom All", input: "</div>", parser: parse.StringFrom(parse.All(parse.String("</"), parse.String("div"), parse.Rune('>')))},
		{name: "StringFrom Times", input: "2026-10-26", parser: parse.StringFrom(parse.Times(4, parse.ZeroToNine), parse.Rune('-'), parse.Times(2, parse.ZeroToNine), parse.Rune('-'), parse.Times(2, parse.ZeroToNine))},
		{name: "StringUntil", input: "some text that runs until the end of a comment -->", parser: parse.StringUntil(parse.String("-->"))},
		{name: "StringUntil RuneNotIn", input: "data-attribute-name=", parser: parse.StringUntil(parse.RuneNotIn("abcdefghijklmnopqrstuvwxyz-"))},
		{name: "Rune", input: ">", parser: parse.Rune('>')},
		{name: "RuneIn", input: "7", parser: parse.ZeroToNine},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			runParserBenchmark(b, bm.input, bm.parser)
		})
	}
}

func BenchmarkParseSlices(b *testing.B) {
	benchmarks := []struct {
		name   string
		input  string
		parser parse.Parser[[]string]
	}{
		{name: "All", input: "</div>", parser: parse.All(parse.String("</"), parse.String("div"), parse.Rune('>'))},
		{name: "Times", input: "12345678", parser: parse.Times(8, parse.ZeroToNine)},
		{name: "ZeroOrMore", input: "aaaaaaaaaaaaaaaab", parser: parse.ZeroOrMore(parse.Rune('a'))},
		{name: "Until", input: "abcdefgh;", parser: parse.Until(parse.Letter, parse.Rune(';'))},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			runParserBenchmark(b, bm.input, bm.parser)
		})
	}
}

func BenchmarkConstructAndParse(b *testing.B) {
	b.ReportAllocs()
	in := parse.NewInput([]byte("</div>"))
	for b.Loop() {
		in.Seek(0)
		_, ok, err := parse.All(parse.String("</"), parse.String("div"), parse.Rune('>')).Parse(in)
		if err != nil || !ok {
			b.Fatalf("failed to parse: ok=%v, err=%v", ok, err)
		}
	}
}

func BenchmarkNewInput(b *testing.B) {
	for _, lines := range []int{10, 1000} {
		data := []byte(strings.Repeat("<p>Some HTML content.</p>\n", lines))
		b.Run(fmt.Sprintf("lines=%d", lines), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				in := parse.NewInput(data)
				in.Seek(len(data) / 2)
				in.Position()
			}
		})
	}
}
