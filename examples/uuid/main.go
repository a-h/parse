package main

import (
	"fmt"

	"github.com/a-h/parse"
)

type UUID string

var hexParser = parse.RuneIn("0123456789abcdefABCDEF")
var uuidStringParser = parse.StringFrom(
	parse.StringFrom(parse.Times(8, hexParser)),
	parse.Rune('-'),
	parse.StringFrom(parse.Times(4, hexParser)),
	parse.Rune('-'),
	parse.StringFrom(parse.Times(4, hexParser)),
	parse.Rune('-'),
	parse.StringFrom(parse.Times(4, hexParser)),
	parse.Rune('-'),
	parse.StringFrom(parse.Times(12, hexParser)),
)
var uuidParser = parse.Convert(uuidStringParser, func(s string) (UUID, error) {
	return UUID(s), nil
})

func main() {
	uuid := "123e4567-e89b-12d3-a456-426655440000"
	input := parse.NewInput(uuid)
	match, ok, err := uuidParser.Parse(input)
	fmt.Println("match:", match)
	fmt.Println("ok:", ok)
	fmt.Println("err:", err)
}
