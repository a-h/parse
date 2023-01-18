package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/a-h/parse"
)

type Identifier string

func createSpaceDelimiterIdentifierParser() parse.Parser[[]Identifier] {
	// Create an identifier parser, letter, followed by letters or numbers.
	// Up until whitespace or the EOF.
	identifier := parse.StringFrom(parse.Letter, parse.StringUntilEOF(parse.Whitespace))
	// Parse multiple identifiers, space delimited.
	// Allow ending on EOF as well as whitespace.
	spaceDelimited := parse.OneOrMore(parse.Then(identifier, parse.Or(parse.Whitespace, parse.EOF[string]())))
	return parse.Func(func(in parse.Input) (match []Identifier, err error) {
		m, err := spaceDelimited.Parse(in)
		if err != nil {
			return
		}
		for _, mm := range m {
			match = append(match, Identifier(mm.A))
		}
		return
	})
}

var SpaceDelimitedIdentifiers = createSpaceDelimiterIdentifierParser()

func main() {
	input := parse.NewInput("validIdentifier1 validIdentifier2")
	match, err := SpaceDelimitedIdentifiers.Parse(input)
	if errors.Is(err, parse.ErrNotMatched) {
		log.Fatal("expected pattern not matched")
	}
	if err != nil {
		log.Fatalf("failed to parse: %v", err)
	}
	fmt.Println(match)
}
