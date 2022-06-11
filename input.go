package parse

func NewInput(s string) *InputString {
	ip := &InputString{
		s:         s,
		charIndex: 0,
	}
	for i, r := range s {
		if r == '\n' {
			ip.newLines = append(ip.newLines, i)
		}
	}
	return ip
}

type InputString struct {
	s         string
	charIndex int
	// character positions of new line characters.
	newLines []int
}

func (in *InputString) Peek(n int) (s string, ok bool) {
	if in.charIndex+n > len(in.s) {
		return
	}
	if n < 0 {
		return in.s[in.charIndex:], true
	}
	return in.s[in.charIndex : in.charIndex+n], true
}

func (in *InputString) Chomp(n int) (s string, ok bool) {
	if in.charIndex+n > len(in.s) {
		return
	}
	from := in.charIndex
	in.charIndex += n
	return in.s[from:in.charIndex], true
}

// Position returns the line and column number of the current position within the stream.
func (in *InputString) Position() (line, column int) {
	var previousLineEnd int
	for lineIndex, lineEnd := range in.newLines {
		if in.charIndex > previousLineEnd && in.charIndex < lineEnd {
			return lineIndex + 1, in.charIndex - previousLineEnd
		}
		previousLineEnd = lineEnd
	}
	return -1, -1
}

// Index returns the current character index of the parser input.
func (in *InputString) Index() int {
	return in.charIndex
}

// Seek to a position in the string.
func (in *InputString) Seek(index int) (ok bool) {
	if index < 0 || index > len(in.s) {
		return
	}
	in.charIndex = index
	return true
}
