package shape

import (
	"io"

	"github.com/gregoryv/nexus"
)

// NewAnchor returns a HTML link entity that wraps any shape.
func NewAnchor(href string, v Shape) *Anchor {
	return &Anchor{
		Href:  href,
		Shape: v,
	}
}

type Anchor struct {
	Href string
	Shape
}

func (a *Anchor) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	w.Printf(`<a href="%s">`, a.Href)
	a.Shape.WriteSVG(w)
	w.Print("</a>")
	return *err
}
