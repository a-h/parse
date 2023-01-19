package parse

import "errors"

type untilParser[T, D any] struct {
	Parser    Parser[T]
	Delimiter Parser[D]
}

// Keep collecting, until we hit the delimiter.
func (p untilParser[T, D]) Parse(in Input) (match []T, err error) {
	for {
		// Attempt to parse something.
		var m T
		m, err = p.Parser.Parse(in)
		if err != nil {
			return
		}
		match = append(match, m)
		// Look for the delimiter.
		start := in.Position()
		_, err = p.Delimiter.Parse(in)
		if err != nil && !errors.Is(err, ErrNotMatched) {
			return
		}
		if err == nil {
			// We found the delimiter, time to quit.
			in.Seek(start.Index)
			err = nil
			return
		}
	}
}

// Until matches until the delimiter is reached.
func Until[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return untilParser[T, D]{
		Parser:    parser,
		Delimiter: delimiter,
	}
}

// UntilEOF matches until the delimiter or the end of the file is reached.
func UntilEOF[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return untilParser[T, D]{
		Parser:    parser,
		Delimiter: Any(delimiter, EOF[D]()),
	}
}
