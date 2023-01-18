package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/a-h/parse"
)

func createYearMonthDayParser() parse.Parser[time.Time] {
	// Create parsers for year, month and day.
	year := parse.StringFrom(parse.Times(4, parse.ZeroToNine))
	month := parse.StringFrom(parse.RuneIn("01"), parse.ZeroToNine)
	day := parse.StringFrom(parse.RuneIn("0123"), parse.ZeroToNine)

	// Create string parser for yyyy-MM-dd.
	// It returns a string array of all the parts.
	date := parse.All(year, parse.Rune('-'), month, parse.Rune('-'), day)

	f := func(in parse.Input) (match time.Time, err error) {
		pos := in.Position()
		var parts []string
		parts, err = date.Parse(in)
		if err != nil {
			return
		}
		var y, m, d int
		y, err = strconv.Atoi(parts[0])
		if err != nil {
			return match, parse.Error(fmt.Sprintf("yearmonthday: invalid year: %v", err), pos)
		}
		m, err = strconv.Atoi(parts[2])
		if err != nil {
			return match, parse.Error(fmt.Sprintf("yearmonthday: invalid month: %v", err), pos)
		}
		d, err = strconv.Atoi(parts[4])
		if err != nil {
			return match, parse.Error(fmt.Sprintf("yearmonthday: invalid day: %v", err), pos)
		}
		match = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		return
	}
	return parse.Func(f)
}

var YearMonthDay = createYearMonthDayParser()

func main() {
	input := parse.NewInput("2000-01-02")
	dateParts, err := YearMonthDay.Parse(input)
	if err != nil {
		log.Fatalf("failed to parse: %v", err)
	}
	fmt.Println(dateParts)
}
