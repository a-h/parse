package parse

type RuneFunctionParser struct {
	F func(r rune) bool
}

func (p RuneFunctionParser) Parse(in Input) (match string, ok bool, err error) {
	match, ok = in.Peek(1)
	if !ok {
		return
	}
	ok = p.F(rune(match[0]))
	if !ok {
		return
	}
	in.Chomp(1)
	return
}

func RuneFunction(f func(r rune) bool) Parser[string] {
	return RuneFunctionParser{
		F: f,
	}
}
