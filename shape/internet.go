package shape

import (
	"io"

	"github.com/gregoryv/draw"
)

func NewInternet() *Internet {
	return &Internet{
		Circle: Circle{Radius: 40, class: "internet"},
		Title:  "Internet",
		Font:   DefaultFont,
		Pad:    DefaultTextPad,
		class:  "internet",
	}
}

type Internet struct {
	Circle
	Title string

	Font  Font
	Pad   Padding
	class string
}

func (r *Internet) WriteSVG(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	r.Circle.WriteSVG(w)
	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *Internet) title() *Label {
	label := &Label{
		Font:  r.Font,
		Text:  r.Title,
		class: "internet-title",
	}
	align := Aligner{}
	align.VAlignCenter(&r.Circle, label)
	align.HAlignCenter(&r.Circle, label)
	Move(label, 0, -18) // todo above HAlignCenter is broken
	return label
}

func (r *Internet) SetFont(f Font)         { r.Font = f }
func (r *Internet) SetTextPad(pad Padding) { r.Pad = pad }
