package parse

import "errors"

type stringUntilParser[T any] struct {
	Delimiter Parser[T]
	AllowEOF  bool
}

func (p stringUntilParser[T]) Parse(in Input) (match string, err error) {
	start := in.Index()
	for {
		beforeDelimiter := in.Index()
		_, err = p.Delimiter.Parse(in)
		if err != nil && !errors.Is(err, ErrNotMatched) {
			return
		}
		if err == nil {
			in.Seek(beforeDelimiter)
			break
		}
		_, chompOK := in.Take(1)
		if !chompOK {
			if p.AllowEOF {
				break
			}
			return "", ErrNotMatched
		}
	}
	end := in.Index()
	in.Seek(start)
	match, _ = in.Take(end - start)
	return match, nil
}

// StringUntil matches until the delimiter is reached.
func StringUntil[T any](delimiter Parser[T]) Parser[string] {
	return stringUntilParser[T]{
		Delimiter: delimiter,
	}
}

// StringUntilEOF matches until the delimiter or the end of the file is reached.
func StringUntilEOF[T any](delimiter Parser[T]) Parser[string] {
	return stringUntilParser[T]{
		Delimiter: delimiter,
		AllowEOF:  true,
	}
}
