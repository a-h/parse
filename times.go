package parse

type TimesParser[T any] struct {
	P Parser[T]
	N int
}

func (p TimesParser[T]) Parse(in Input) (match []T, ok bool, err error) {
	start := in.Index()
	match = make([]T, p.N)
	for i := 0; i < p.N; i++ {
		var m T
		m, ok, err = p.P.Parse(in)
		if err != nil {
			return match, false, err
		}
		if !ok {
			in.Seek(start)
			return nil, false, nil
		}
		if ok {
			match[i] = m
		}
	}
	return match, true, nil
}

func Times[T any](n int, p Parser[T]) TimesParser[T] {
	return TimesParser[T]{
		P: p,
		N: n,
	}
}
