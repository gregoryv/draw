package design

import "io"

type ClassDiagram struct {
	Diagram
}

func NewClassDiagram() *ClassDiagram {
	return &ClassDiagram{}
}

func (d *ClassDiagram) WriteSvg(w io.Writer) error {
	return nil
}
