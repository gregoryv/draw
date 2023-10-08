package shape

import (
	"html/template"

	"github.com/gregoryv/draw"
)

// NewLabeled returns a shape with the text as a label below it.
func NewLabeled(text string, s Shape) *Labeled {
	l := &Label{
		text:  template.HTMLEscapeString(text),
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultPad,
		class: "label",
	}
	// todo place below
	h := s.Height()
	l.SetY(h)

	var a Aligner
	sw := s.Width()
	lw := l.Width()
	if sw > lw {
		a.VAlignCenter(s, l)
	} else {
		a.VAlignCenter(l, s)
	}
	return &Labeled{NewGroup(l, s)}
}

type Labeled struct {
	*Group
}
