package parse

import "fmt"

type Position struct {
	Index, Line, Col int
}

func (pos Position) String() string {
	// The line number is 0-based, but text editors are usually 1-based.
	// So the number is incremented by 1 for display.
	return fmt.Sprintf("line %d, col %d", pos.Line+1, pos.Col)
}

type pf[T any] struct {
	f func(in *Input) (item T, ok bool, err error)
}

func (p pf[T]) Parse(in *Input) (item T, ok bool, err error) {
	return p.f(in)
}

func (p pf[T]) Match(in *Input) (ok bool, err error) {
	_, ok, err = p.f(in)
	return
}

// Func creates a parser from an input function.
func Func[T any](f func(in *Input) (item T, ok bool, err error)) Parser[T] {
	return pf[T]{
		f: f,
	}
}

// Matcher matches the input and advances past the match without building a
// value from it.
type Matcher interface {
	Match(in *Input) (ok bool, err error)
}

// Parser matches the input and builds a value from the match. Match must
// return the same ok and err, and leave the input at the same position, as
// Parse.
type Parser[T any] interface {
	Matcher
	Parse(in *Input) (item T, ok bool, err error)
}

// Match parses the input with p and discards the value. Use it to implement
// the Match method of a parser that has no cheaper way to match the input.
func Match[P Parser[T], T any](p P, in *Input) (ok bool, err error) {
	_, ok, err = p.Parse(in)
	return
}
