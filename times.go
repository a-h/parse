package parse

import "math"

type timesParser[T any] struct {
	P   Parser[T]
	Min int
	Max int
}

func (p timesParser[T]) Parse(in *Input) (match []T, ok bool, err error) {
	start := in.Index()
	for len(match) < p.Max {
		var m T
		m, ok, err = p.P.Parse(in)
		if err != nil {
			return match, false, err
		}
		if !ok {
			break
		}
		if match == nil {
			match = make([]T, 0, max(p.Min, 1))
		}
		match = append(match, m)
	}
	ok = len(match) >= p.Min && len(match)-1 < p.Max
	if !ok {
		in.Seek(start)
		return match, false, nil
	}
	return match, true, nil
}

func (p timesParser[T]) Match(in *Input) (ok bool, err error) {
	start := in.Index()
	var count int
	for count < p.Max {
		ok, err = p.P.Match(in)
		if err != nil {
			return false, err
		}
		if !ok {
			break
		}
		count++
	}
	ok = count >= p.Min && count-1 < p.Max
	if !ok {
		in.Seek(start)
		return false, nil
	}
	return true, nil
}

// Times matches the given parser n times.
func Times[T any](n int, p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: n,
		Max: n,
	}
}

// Repeat matches the given parser between min and max times.
func Repeat[T any](min, max int, p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: min,
		Max: max,
	}
}

// AtLeast matches the given parser at least min times.
func AtLeast[T any](min int, p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: min,
		Max: math.MaxInt,
	}
}

// AtMost matches the given parser at most max times.
// It is equivalent to ZeroOrMore.
func AtMost[T any](max int, p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: 0,
		Max: max,
	}
}

// ZeroOrMore matches the given parser zero or more times.
func ZeroOrMore[T any](p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: 0,
		Max: math.MaxInt,
	}
}

// OneOrMore matches the given parser at least once.
func OneOrMore[T any](p Parser[T]) Parser[[]T] {
	return timesParser[T]{
		P:   p,
		Min: 1,
		Max: math.MaxInt,
	}
}
