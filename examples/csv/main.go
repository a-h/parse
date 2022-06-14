package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/a-h/parse"
)

var colDelimiter = parse.Rune(',')
var doubleQuote = parse.Rune('"')
var escape = parse.Rune('\\')
var escapedQuote = parse.String(`\"`)
var stringUntilEscapedCharacterOrDoubleQuote = parse.StringUntil(parse.Any(escape, doubleQuote))

type QuotedStringParser struct{}

func (p QuotedStringParser) Parse(in parse.Input) (match string, ok bool, err error) {
	start := in.Index()
	// Start with a quote.
	_, ok, err = doubleQuote.Parse(in)
	if err != nil || !ok {
		// No match, so rewind.
		in.Seek(start)
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
		err = fmt.Errorf("unterminated quoted string from %v to %v", in.PositionAt(start), in.Position())
	}
	match = sb.String()
	return
}

var quotedString = parse.Parser[string](QuotedStringParser{})
var rowDelimiter = parse.NewLine
var unquotedString = parse.StringUntil(parse.Any(colDelimiter, rowDelimiter, parse.EOF[string]()))
var stringValueParser = parse.Func(func(in parse.Input) (match string, ok bool, err error) {
	match, ok, err = parse.Any(quotedString, unquotedString).Parse(in)
	// Chomp the col delimiter, but we could also be at the end of the row, or file.
	parse.Any(colDelimiter, rowDelimiter, parse.EOF[string]()).Parse(in)
	return
})

var row parse.Parser[[]string] = parse.Func(func(in parse.Input) (match []string, ok bool, err error) {
	match, ok, err = parse.UntilEOF(stringValueParser, rowDelimiter).Parse(in)
	// Chomp the row terminator.
	rowDelimiter.Parse(in)
	return
})

var CSV parse.Parser[[][]string] = parse.UntilEOF(row, parse.Any(colDelimiter, parse.NewLine))

func main() {
	csvData := `a,b,c
"d",e,f
"a string",123,456`
	input := parse.NewInput(csvData)
	//match, ok, err := CSV.Parse(input)
	match, ok, err := row.Parse(input)
	//match, ok, err := stringValueParser.Parse(input)
	if err != nil {
		log.Fatalf("failed to parse: %v", err)
	}
	if !ok {
		log.Print("expected CSV data not matched")
	}
	j, _ := json.Marshal(match)
	fmt.Println(string(j))
}
