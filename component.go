package design

import (
	"io"

	"github.com/gregoryv/go-design/xml"
)

type Component struct {
	Label string
}

func (comp *Component) WriteTo(w io.Writer) (int, error) {
	all := make(Drawables, 0)
	all = append(all, xml.NewNode(Element_rect,
		x("30"), y("20"), width("250"), height("150"),
		style("fill:#ffffcc;stroke:black;stroke-width:1;opacity:0.5"),
	))
	text := xml.NewNode(Element_text,
		x("50"), y("55"), fill("black"),
	)
	text.Append(xml.CData(comp.Label))
	all = append(all, text)
	return all.WriteTo(w)
}
