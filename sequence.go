package design

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/shape"
)

type SequenceDiagram struct {
	Width, Height int
	ColWidth      int
	Font          shape.Font
	TextPad       shape.Padding
	Pad           shape.Padding

	columns []string
	links   []link
}

func (dia *SequenceDiagram) WriteSvg(w io.Writer) error {
	svg := &shape.Svg{
		Width:  dia.Width,
		Height: dia.Height,
	}

	colWidth := dia.ColWidth
	top := dia.Font.Height + dia.TextPad.Top + dia.Pad.Top
	// todo use correct padding around diagram
	var (
		x  = colWidth / 2
		y1 = top + dia.TextPad.Bottom // below label
		y2 = dia.Height
	)
	lines := make([]*shape.Line, len(dia.columns))
	for i, column := range dia.columns {
		lines[i] = &shape.Line{Class: "columnline", X1: x, Y1: y1, X2: x, Y2: y2}
		x += colWidth

		label := &shape.Label{X: i * colWidth, Y: top,
			Text: column,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		shape.AlignVertical(shape.Center, lines[i], label)
		svg.Content = append(svg.Content, lines[i], label)
	}

	for i, link := range dia.links {
		k := i + 1
		y := y1 + (dia.Font.LineHeight+dia.Pad.Bottom)*k
		arrow := &shape.Arrow{}
		arrow.X1 = lines[link.fromIndex].X1
		arrow.Y1 = y
		arrow.X2 = lines[link.toIndex].X1
		arrow.Y2 = y

		label := &shape.Label{
			X:    arrow.X1,
			Y:    y - 2,
			Text: link.text,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		shape.AlignVertical(shape.Center, arrow, label)
		svg.Content = append(svg.Content, arrow, label)
	}
	return svg.WriteSvg(w)
}

func (dia *SequenceDiagram) AddColumns(names ...string) {
	dia.columns = append(dia.columns, names...)
}

func (dia *SequenceDiagram) Link(from, to, text string) error {
	fromIndex := -1
	toIndex := -1
	for i, column := range dia.columns {
		if column == from {
			fromIndex = i
			continue
		}
		if column == to {
			toIndex = i
		}
	}
	link := link{fromIndex, toIndex, text}
	dia.links = append(dia.links, link)
	if fromIndex == -1 {
		return fmt.Errorf("Missing %q column", from)
	}
	if toIndex == -1 {
		return fmt.Errorf("Missing %q column", to)
	}
	return nil
}

type link struct {
	fromIndex, toIndex int
	text               string
}
