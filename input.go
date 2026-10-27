package parse

import (
	"bytes"
	"slices"
	"unsafe"
)

// NewInput creates an input from b. The input takes ownership of b and reads
// it in place without copying it. Strings returned by Peek and Take share
// memory with b, so the caller must not modify b after calling NewInput.
func NewInput(b []byte) *Input {
	ip := &Input{
		b:         b,
		s:         unsafe.String(unsafe.SliceData(b), len(b)),
		charIndex: 0,
		newLines:  make([]int, 0, bytes.Count(b, []byte{'\n'})),
	}
	var offset int
	for {
		i := bytes.IndexByte(b[offset:], '\n')
		if i < 0 {
			break
		}
		ip.newLines = append(ip.newLines, offset+i)
		offset += i + 1
	}
	return ip
}

// InputString is an input used by parsers. It stores the current location
// and character positions.
type Input struct {
	b []byte
	// s is a string view of b that shares its memory.
	s         string
	charIndex int
	// character positions of new line characters.
	newLines []int
}

// Peek returns the next n bytes as a string without changing the position.
// If n is negative, it returns the rest of the input.
func (in *Input) Peek(n int) (s string, ok bool) {
	if in.charIndex+n > len(in.s) {
		return
	}
	if n < 0 {
		return in.s[in.charIndex:], true
	}
	return in.s[in.charIndex : in.charIndex+n], true
}

// PeekBytes returns the next n bytes without changing the position. If n is
// negative, it returns the rest of the input. The slice shares memory with the
// input, so callers must not modify it.
func (in *Input) PeekBytes(n int) (b []byte, ok bool) {
	if in.charIndex+n > len(in.b) {
		return
	}
	if n < 0 {
		return in.b[in.charIndex:len(in.b):len(in.b)], true
	}
	return in.b[in.charIndex : in.charIndex+n : in.charIndex+n], true
}

// Take returns the next n bytes as a string and moves the position past them.
func (in *Input) Take(n int) (s string, ok bool) {
	if in.charIndex+n > len(in.b) {
		return
	}
	from := in.charIndex
	in.charIndex += n
	return in.s[from:in.charIndex], true
}

// Position returns the zero-bound index, line and column number of the current position within the stream.
func (in *Input) Position() Position {
	return in.PositionAt(in.charIndex)
}

// Position returns the zero-bound index, line and column number of the current position within the stream.
func (in *Input) PositionAt(index int) Position {
	lineIndex, _ := slices.BinarySearch(in.newLines, index)
	var previousLineEnd int
	if lineIndex > 0 {
		previousLineEnd = in.newLines[lineIndex-1] + 1
	}
	colIndex := index - previousLineEnd
	return Position{Index: index, Line: lineIndex, Col: colIndex}
}

// Index returns the current character index of the parser input.
func (in *Input) Index() int {
	return in.charIndex
}

// Seek to a position in the string.
func (in *Input) Seek(index int) (ok bool) {
	if index < 0 || index > len(in.b) {
		return
	}
	in.charIndex = index
	return true
}
