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
		Parts: make(Drawables, 0),
	}
}

type Graph struct {
	Width, Height int
	Title         string
	Parts         Drawables
}

func (graph *Graph) String() string {
	buf := bytes.NewBufferString("")
	graph.Parts.WriteTo(buf)
	return buf.String()
}

func (graph *Graph) NewComponent(v interface{}) {
	component := &Component{
		Label: reflect.TypeOf(v).Name(),
	}
	graph.Title = component.Label
	graph.Parts = append(graph.Parts, component)
}

const header string = `<svg width="{{.Width}}" height="{{.Height}}"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<title>{{.Title}}</title>
`

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

type Component struct {
	Label string
}

func (comp *Component) WriteTo(w io.Writer) (int, error) {
	all := make(Drawables, 0)
	all = append(all, NewNode(Element_rect,
		x("30"), y("20"), width("150"), height("150"),
		style("fill:#ffffcc;stroke:black;stroke-width:1;opacity:0.5"),
	))
	return all.WriteTo(w)
	/*
	   <text x="50" y="55" fill="black">Account</text>  */
}

func x(v string) Attribute      { return Attribute{"x", v} }
func y(v string) Attribute      { return Attribute{"y", v} }
func width(v string) Attribute  { return Attribute{"width", v} }
func height(v string) Attribute { return Attribute{"width", v} }
func style(v string) Attribute  { return Attribute{"style", v} }
