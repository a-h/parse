package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInputString(t *testing.T) {
	s := `Line 0
Line 1
Line 2
Line 3`

	in := NewInput(s)
	t.Run("can peek forwards", func(t *testing.T) {
		in.Seek(0)
		r, ok := in.Peek(2)
		if !ok {
			t.Errorf("expected OK, got %v", ok)
		}
		if diff := cmp.Diff("Li", r); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("peeking doesn't change position", func(t *testing.T) {
		in.Seek(0)
		in.Peek(2)
		want := Position{0, 0, 0}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("chomping up to a line ending changes position, but doesn't start a newline", func(t *testing.T) {
		in.Seek(0)
		in.Take(len("Line 1"))
		want := Position{6, 0, 6}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("chomping a line ending changes position", func(t *testing.T) {
		in.Seek(0)
		in.Take(len("Line 1\n"))
		want := Position{7, 1, 0}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("but you can seek to any point", func(t *testing.T) {
		in.Seek(5)
		want := Position{5, 0, 5}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("partial lines can be read", func(t *testing.T) {
		in.Seek(7)
		s, ok := in.Take(len("Line "))
		if !ok {
			t.Fatalf("failed to chomp line")
		}
		if diff := cmp.Diff("Line ", s); diff != "" {
			t.Error(diff)
		}
		want := Position{12, 1, 5}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("all the lines can be read", func(t *testing.T) {
		in.Seek(0)
		in.Take(len("Line 1\n"))
		in.Take(len("Line 2\n"))
		in.Take(len("Line 3\n"))
		want := Position{21, 3, 0}
		got := in.Position()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
}

func TestPositionAt(t *testing.T) {
	input := `Here's some
test text, just
to test the
newline count.`
	pi := NewInput(input)
	var positions []Position
	for {
		positions = append(positions, pi.Position())
		if _, ok := pi.Take(1); !ok {
			break
		}
	}
	for i := 0; i < len(positions)-1; i++ {
		if diff := cmp.Diff(positions[i], pi.PositionAt(i)); diff != "" {
			t.Errorf("position %d: %s", i, diff)
		}
	}
}
