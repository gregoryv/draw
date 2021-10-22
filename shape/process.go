package shape

import (
	"io"
)

func NewProcess(text string) *Process {
	label := NewLabel(text)
	label.SetClass("process-title")
	return &Process{
		Circle: Circle{
			Radius: (label.Width() + DefaultTextPad.Left + DefaultTextPad.Right) / 2,
			class:  "process",
		},
		label: label,
	}
}

type Process struct {
	Circle
	label *Label
}

func (r *Process) WriteSVG(out io.Writer) error {
	r.Circle.WriteSVG(out)
	var a Aligner
	a.HAlignCenter(&r.Circle, r.label)
	a.VAlignCenter(&r.Circle, r.label)
	return r.label.WriteSVG(out)
}
