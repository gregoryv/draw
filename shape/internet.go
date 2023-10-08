package shape

import (
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/nexus"
)

func NewInternet() *Internet {
	return &Internet{
		Circle: Circle{Radius: 40, class: "internet"},
		Title:  "Internet",
		Font:   draw.DefaultFont,
		Pad:    draw.DefaultTextPad,
		class:  "internet",
	}
}

type Internet struct {
	Circle
	Title string

	Font  draw.Font
	Pad   draw.Padding
	class string
}

func (r *Internet) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	r.Circle.WriteSVG(w)
	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *Internet) title() *Label {
	label := &Label{
		Font:  r.Font,
		text:  r.Title,
		class: "internet-title",
	}
	label.Font.LineHeight = draw.DefaultFont.Height
	align := Aligner{}
	align.VAlignCenter(&r.Circle, label)
	align.HAlignCenter(&r.Circle, label)
	return label
}

func (r *Internet) SetFont(f draw.Font)         { r.Font = f }
func (r *Internet) SetTextPad(pad draw.Padding) { r.Pad = pad }
