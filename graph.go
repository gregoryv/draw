package design

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"reflect"
)

func NewGraph() *Graph {
	return &Graph{
		Width:  100, // Default size, should be adapted by content I think
		Height: 100,
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

func (graph *Graph) NewComponent(v interface{}) {
	component := &Component{
		Label: reflect.TypeOf(v).Name(),
	}
	graph.Title = component.Label
	graph.Parts = append(graph.Parts, component)
}
