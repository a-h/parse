package parse

type eofParser struct {
}

func (p eofParser) Parse(in Input) (match any, ok bool, err error) {
	_, canAdvance := in.Peek(1)
	ok = !canAdvance
	return
}

// EOF matches the end of the input.
var EOF Parser[any] = eofParser{}
