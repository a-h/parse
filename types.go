package parse

// Input represents the input to a parser.
type Input interface {
	// Peek n runes ahead, returns !ok if it's not possible to read that much.
	// If n < 0, the remainder of the input is returned.
	Peek(n int) (s string, ok bool)
	// Advance by a number of runes, returns !ok if it's no possible because the end has been reached.
	Chomp(n int) (s string, ok bool)
	// Position returns the line and column number of the current position.
	Position() (line, column int)
	// Index returns the current rune index of the parser input.
	Index() int
	// Seek to a location in the input.
	Seek(index int) (ok bool)
}

// Parser is implemented by all parsers.
type Parser[T any] interface {
	Parse(in Input) (item T, ok bool, err error)
}
