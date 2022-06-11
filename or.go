package parse

type OrParser[A any, B any] struct {
	A Parser[A]
	B Parser[B]
}

type Option[T any] struct {
	Value T
	OK    bool
}

func OptionValue[T any](v T, ok bool) Option[T] {
	return Option[T]{
		Value: v,
		OK:    ok,
	}
}

type OrResult[A any, B any] struct {
	A Option[A]
	B Option[B]
}

func (p OrParser[A, B]) Parse(in Input) (match OrResult[A, B], ok bool, err error) {
	a, ok, err := p.A.Parse(in)
	if err != nil {
		return
	}
	if ok {
		return OrResult[A, B]{A: OptionValue(a, true)}, ok, err
	}
	b, ok, err := p.B.Parse(in)
	if err != nil {
		return
	}
	if ok {
		return OrResult[A, B]{B: OptionValue(b, true)}, ok, err
	}
	return
}

func Or[A any, B any](a Parser[A], b Parser[B]) OrParser[A, B] {
	return OrParser[A, B]{
		A: a,
		B: b,
	}
}
