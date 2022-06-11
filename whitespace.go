package parse

import "unicode"

// Whitespace parses whitespace.
var Whitespace Parser[string] = StringFrom(OneOrMore(RuneInRanges(unicode.White_Space)))

// OptionalWhitespace parses optional whitespace.
var OptionalWhitespace Parser[OptionalResult[string]] = Optional(Whitespace)
