package parse_test

import (
	"fmt"
	"strconv"

	"github.com/a-h/parse"
)

func ExampleString() {
	abParser := parse.Any(parse.String("A"))

	fmt.Println(abParser.Parse(parse.NewInput("A")))
	fmt.Println(abParser.Parse(parse.NewInput("B")))
	// Output:
	// A true <nil>
	//  false <nil>
}

func ExampleAny() {
	abParser := parse.Any(parse.String("A"), parse.String("B"))

	fmt.Println(abParser.Parse(parse.NewInput("A")))
	fmt.Println(abParser.Parse(parse.NewInput("B")))
	fmt.Println(abParser.Parse(parse.NewInput("C")))
	// Output:
	// A true <nil>
	// B true <nil>
	//  false <nil>
}

func ExampleAll() {
	abcParser := parse.All(parse.String("A"), parse.String("B"), parse.String("C"))

	fmt.Println(abcParser.Parse(parse.NewInput("ABC")))
	fmt.Println(abcParser.Parse(parse.NewInput("AB")))
	fmt.Println(abcParser.Parse(parse.NewInput("A")))
	// Output:
	// [A B C] true <nil>
	// [A B] false <nil>
	// [A] false <nil>
}

func ExampleOptional() {
	abcParser := parse.StringFrom(
		parse.StringFrom(parse.Optional(parse.String("A"))),
		parse.String("B"),
	)

	fmt.Println(abcParser.Parse(parse.NewInput("ABC")))
	fmt.Println(abcParser.Parse(parse.NewInput("B")))
	fmt.Println(abcParser.Parse(parse.NewInput("A")))
	// Output:
	// AB true <nil>
	// B true <nil>
	//  false <nil>
}

func ExampleParser() {
	type GotoStatement struct {
		Line int64
	}
	gotoParser := parse.Func(func(in *parse.Input) (item GotoStatement, ok bool, err error) {
		start := in.Index()

		if _, ok, err = parse.String("GOTO ").Parse(in); err != nil || !ok {
			// Rollback, and return.
			in.Seek(start)
			return
		}

		// Read until the next newline or the EOF.
		until := parse.Any(parse.NewLine, parse.EOF[string]())
		var lineNumber string
		if lineNumber, ok, err = parse.StringUntil(until).Parse(in); err != nil || !ok {
			err = parse.Error("Syntax error: GOTO is missing line number", in.Position())
			return
		}
		// We must have a valid line number now, or there is a syntax error.
		item.Line, err = strconv.ParseInt(lineNumber, 10, 64)
		if err != nil {
			return item, false, parse.Error("Syntax error: GOTO has invalid line number", in.Position())
		}

		// Chomp the newline we read up to.
		until.Parse(in)

		return item, true, nil
	})

	inputs := []string{
		"GOTO 10",
		"GOTO abc",
		"FOR i = 0",
	}

	for _, input := range inputs {
		stmt, ok, err := gotoParser.Parse(parse.NewInput(input))
		fmt.Printf("%+v, ok=%v, err=%v\n", stmt, ok, err)
	}
	// Output:
	// {Line:10}, ok=true, err=<nil>
	// {Line:0}, ok=false, err=Syntax error: GOTO has invalid line number: line 0, col 8
	// {Line:0}, ok=false, err=<nil>
}
