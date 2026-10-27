package parse_test

import (
	"errors"

	"github.com/a-h/parse"
)

var errTestParseError = errors.New("parser failure")

type expectErrorParser struct {
}

func (p expectErrorParser) Parse(in *parse.Input) (match string, ok bool, err error) {
	err = errTestParseError
	return
}

func (p expectErrorParser) Match(in *parse.Input) (ok bool, err error) {
	return parse.Match(p, in)
}
