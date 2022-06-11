package parse

type TimesParser[T any] struct {
	P   Parser[T]
	Min int
	Max func(i int) int
}

func (p TimesParser[T]) Parse(in Input) (match []T, ok bool, err error) {
	start := in.Index()
	for i := 0; i < p.Max(i); i++ {
		var m T
		m, ok, err = p.P.Parse(in)
		if err != nil {
			return match, false, err
		}
		if !ok {
			break
		}
		if ok {
			match = append(match, m)
		}
	}
	ok = len(match) >= p.Min && len(match) <= p.Max(len(match))
	if !ok {
		in.Seek(start)
		return nil, false, nil
	}
	return match, true, nil
}

func Times[T any](n int, p Parser[T]) Parser[[]T] {
	return TimesParser[T]{
		P:   p,
		Min: n,
		Max: func(i int) int { return n },
	}
}

func Repeat[T any](min, max int, p Parser[T]) Parser[[]T] {
	return TimesParser[T]{
		P:   p,
		Min: min,
		Max: func(i int) int { return max },
	}
}

func AtLeast[T any](min int, p Parser[T]) Parser[[]T] {
	return TimesParser[T]{
		P:   p,
		Min: min,
		Max: func(i int) int { return i + 1 },
	}
}

func AtMost[T any](max int, p Parser[T]) Parser[[]T] {
	return TimesParser[T]{
		P:   p,
		Min: 0,
		Max: func(i int) int { return max },
	}
}
