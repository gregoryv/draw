package design

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/gregoryv/go-design/svg"
)

func NewGraph() *Graph {
	return &Graph{
		Width:  800, // Default size, should be adapted by content I think
		Height: 400,
		Parts:  make(Drawables, 0),
	}
}

type Graph struct {
	Width, Height int
	Title         string
	Parts         Drawables
}

func (graph *Graph) String() string {
	buf := bytes.NewBufferString("")
	graph.WriteTo(buf)
	return buf.String()
}

const header string = `<svg width="{{.Width}}" height="{{.Height}}"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<title>{{.Title}}</title>
`

// WriteTo includes full xml
func (graph *Graph) WriteTo(w io.Writer) {
	tpl := template.Must(template.New("header").Parse(header))
	tpl.Execute(w, graph)

	graph.Parts.WriteTo(w)
	fmt.Fprint(w, "\n</svg>")
}

func (graph *Graph) Link(from, to *Component) {
	if !from.areLinked(to) {
		panic(fmt.Sprintf("Cannot link %v with %v", from.v, to.v))
	}
	x1, y1 := from.Center()
	x2, y2 := to.Center()
	graph.Parts = append(
		Drawables{
			svg.Line(x1, y1, x2, y2),
		},
		graph.Parts...,
	)
}

type Drawable interface {
	WriteTo(io.Writer) (int, error)
}

type Drawables []Drawable

func (all Drawables) WriteTo(w io.Writer) (int, error) {
	var total int
	for _, part := range all {
		n, err := part.WriteTo(w)
		if err != nil {
			return total + n, err
		}
		total += n
	}
	return total, nil
}

func (graph *Graph) Add(d ...Drawable) {
	graph.Parts = append(graph.Parts, d...)
}

func (graph *Graph) Place(obj PositionedDrawable, x, y int) {
	obj.SetX(x)
	obj.SetY(y)
	graph.Add(obj)
}
