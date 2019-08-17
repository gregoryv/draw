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
		x  = dia.Pad.Left
		y1 = top + dia.TextPad.Bottom // below label
		y2 = dia.Height
	)
	lines := make([]*shape.Line, len(dia.columns))
	for i, column := range dia.columns {
		label := &shape.Label{X: i * colWidth, Y: top,
			Text: column,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		if i == 0 {
			x += label.Width() / 2
		}
		lines[i] = &shape.Line{Class: "columnline", X1: x, Y1: y1, X2: x, Y2: y2}
		x += colWidth

		shape.AlignVertical(shape.Center, lines[i], label)
		svg.Content = append(svg.Content, lines[i], label)
	}

	y := y1 + (dia.Font.LineHeight + dia.Pad.Bottom)
	for _, link := range dia.links {
		fromX := lines[link.fromIndex].X1
		toX := lines[link.toIndex].X1
		label := &shape.Label{
			X:    fromX,
			Y:    y - 2,
			Text: link.text,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		arrow := &shape.Arrow{}
		if link.toSelf() {
			margin := 15
			// add two lines + arrow
			l1 := &shape.Line{Class: "arrow", X1: fromX, Y1: y, X2: fromX + margin, Y2: y}
			l2 := &shape.Line{Class: "arrow",
				X1: fromX + margin,
				Y1: y,
				X2: fromX + margin,
				Y2: y + dia.Font.LineHeight*2,
			}
			shape.AlignHorizontal(shape.Center, l2, label)
			label.X += l1.Width() + dia.TextPad.Left
			arrow.X1 = l2.X2
			arrow.Y1 = l2.Y2
			arrow.X2 = l1.X1
			arrow.Y2 = l2.Y2
			svg.Content = append(svg.Content, l1, l2, arrow, label)
			y += 3*dia.Font.LineHeight + dia.Pad.Bottom
		} else {
			arrow.X1 = fromX
			arrow.Y1 = y
			arrow.X2 = toX
			arrow.Y2 = y
			shape.AlignVertical(shape.Center, arrow, label)
			svg.Content = append(svg.Content, arrow, label)
			y += dia.Font.LineHeight + dia.Pad.Bottom
		}

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
			break
		}
	}
	for i, column := range dia.columns {
		if column == to {
			toIndex = i
			break
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

func (l *link) toSelf() bool {
	return l.fromIndex == l.toIndex
}
