package design

import (
	"html/template"
	"io"
)

func NewGraph() *Graph {
	return &Graph{
		Elements: make([]*Element, 0),
	}
}

type Graph struct {
	Width, Height int
	Title         string
	Elements      []*Element
}

func (graph *Graph) WriteTo(w io.Writer) {
	tpl := template.Must(template.New("svg").Parse(svgSource))
	tpl.Execute(w, graph)
}

func (graph *Graph) NewFolder(basename string) {
}

type Element struct{}
