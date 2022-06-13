package parse

// Input represents the input to a parser.
type Input interface {
	// Peek n runes ahead, returns !ok if it's not possible to read that much.
	// If n < 0, the remainder of the input is returned.
	Peek(n int) (s string, ok bool)
	// Advance by a number of runes, returns !ok if it's no possible because the end has been reached.
	Chomp(n int) (s string, ok bool)
	// Position returns the zero-bound index, line and column number of the current position within the stream.
	Position() Position
	// Index returns the current rune index of the parser input.
	Index() int
	// Seek to a location in the input.
	Seek(index int) (ok bool)
}

type Position struct {
	Index, Line, Col int
}

type pf[T any] struct {
	f func(in Input) (item T, ok bool, err error)
}

func (p pf[T]) Parse(in Input) (item T, ok bool, err error) {
	return p.f(in)
}

// Func creates a parser from an input function.
func Func[T any](f func(in Input) (item T, ok bool, err error)) Parser[T] {
	return pf[T]{
		f: f,
	}
}

// Parser is implemented by all parsers.
type Parser[T any] interface {
	Parse(in Input) (item T, ok bool, err error)
}
