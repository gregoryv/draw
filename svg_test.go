package draw

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
)

func ExampleNewSvg() {
	s := NewSvg()
	s.WriteSVG(os.Stdout)
	// output:
	// <svg
	//   xmlns="http://www.w3.org/2000/svg"
	//   xmlns:xlink="http://www.w3.org/1999/xlink"
	//   width="100" height="100" font-family="Arial, Helvetica, sans-serif"></svg>
}

func TestSvg_SetSize(t *testing.T) {
	s := &SVG{}
	s.SetSize(1, 2)
	assert := asserter.New(t)
	assert().Equals(s.Width(), 1)
	assert().Equals(s.Height(), 2)
}

func TestSvg_empty_by_default(t *testing.T) {
	s := NewSvg()
	if len(s.Content) != 0 {
		t.Error("Not empty")
	}
}

func TestSvg_Append(t *testing.T) {
	s := NewSvg()
	shape := &dummy{}
	s.Append(shape)
	if s.Content[0] != shape {
		t.Error("Not first")
	}

	s.Append(shape)
	if s.Content[len(s.Content)-1] != shape {
		t.Error("Not last")
	}
}

func TestPrepend(t *testing.T) {
	s := NewSvg()
	shape := &dummy{}
	s.Prepend(shape)
	if s.Content[0] != shape {
		t.Error("Not first")
	}
}

func TestSvg_WriteSvg(t *testing.T) {
	s := NewSvg()
	shape := &dummy{}
	s.Append(shape)
	w := bytes.NewBufferString("")
	s.WriteSVG(w)
	if w.String() == "" {
		t.Error("No svg written")
	}
}

type dummy struct{}

func (d *dummy) WriteSVG(w io.Writer) error {
	_, err := w.Write([]byte("..."))
	return err
}
