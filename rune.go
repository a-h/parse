package parse

func Rune(r rune) Parser[string] {
	return StringParser{
		Match: string(r),
	}
}
