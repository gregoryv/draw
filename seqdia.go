package design

import (
	"io"

	"github.com/gregoryv/go-design/shape"
)

func NewSequenceDiagram() *SequenceDiagram {
	return &SequenceDiagram{
		ColWidth: 190,
		Diagram:  NewDiagram(),
	}
}

type SequenceDiagram struct {
	Diagram
	ColWidth int

	columns []string
	links   []*Link
}

func (dia *SequenceDiagram) WriteSvg(w io.Writer) error {
	svg := &shape.Svg{
		Width:  dia.Width(),
		Height: dia.Height(),
	}
	var (
		colWidth = dia.ColWidth

		top = dia.top()
		x   = dia.Pad.Left
		y1  = top + dia.TextPad.Bottom // below label
		y2  = dia.Height()
	)
	lines := make([]*shape.Line, len(dia.columns))
	for i, column := range dia.columns {
		label := &shape.Label{X: i * colWidth, Y: top,
			Text: column,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		firstColumn := i == 0
		if firstColumn {
			x += label.Width() / 2
		}
		line := shape.NewLine(x, y1, x, y2)
		line.Class = "column-line"
		lines[i] = line
		x += colWidth

		dia.VAlignCenter(lines[i], label)
		svg.Content = append(svg.Content, lines[i], label)
	}

	y := y1 + dia.plainHeight()
	for _, lnk := range dia.links {
		fromX := lines[lnk.fromIndex].Start.X
		toX := lines[lnk.toIndex].Start.X
		label := &shape.Label{
			X:    fromX,
			Y:    y - 2,
			Text: lnk.text,
			Font: dia.Font,
			Pad:  dia.Pad,
		}
		arrow := &shape.Arrow{}
		if lnk.toSelf() {
			margin := 15
			// add two lines + arrow
			l1 := shape.NewLine(fromX, y, fromX+margin, y)
			l1.Class = lnk.class()
			l2 := shape.NewLine(fromX+margin, y, fromX+margin, y+dia.Font.LineHeight*2)
			l2.Class = lnk.class()
			dia.HAlignCenter(l2, label)
			label.X += l1.Width() + dia.TextPad.Left
			arrow.Start.X = l2.End.X
			arrow.Start.Y = l2.End.Y
			arrow.End.X = l1.Start.X
			arrow.End.Y = l2.End.Y
			arrow.Class = lnk.class()
			svg.Content = append(svg.Content, l1, l2, arrow, label)
			y += dia.selfHeight()
		} else {
			arrow.Start.X = fromX
			arrow.Start.Y = y
			arrow.End.X = toX
			arrow.End.Y = y
			dia.VAlignCenter(arrow, label)
			svg.Content = append(svg.Content, arrow, label)
			y += dia.plainHeight()
		}
	}
	return svg.WriteSvg(w)
}

// Width returns the total width of the diagram
func (dia *SequenceDiagram) Width() int {
	if dia.Svg.Width != 0 {
		return dia.Svg.Width
	}
	return len(dia.columns) * dia.ColWidth
}

// Height returns the total height of the diagram
func (dia *SequenceDiagram) Height() int {
	if dia.Svg.Height != 0 {
		return dia.Svg.Height
	}
	if len(dia.columns) == 0 {
		return 0
	}
	height := dia.top() + dia.plainHeight()
	for _, lnk := range dia.links {
		if lnk.toSelf() {
			height += dia.selfHeight()
			continue
		}
		height += dia.plainHeight()
	}
	return height
}

// selfHeight is the height of a self referencing link
func (dia *SequenceDiagram) selfHeight() int {
	return 3*dia.Font.LineHeight + dia.Pad.Bottom
}

// plainHeight returns the height of and arrow and label
func (dia *SequenceDiagram) plainHeight() int {
	return dia.Font.LineHeight + dia.Pad.Bottom
}

func (dia *SequenceDiagram) top() int {
	return dia.Font.LineHeight + dia.Pad.Top
}

func (dia *SequenceDiagram) AddColumns(names ...string) {
	dia.columns = append(dia.columns, names...)
}

func (dia *SequenceDiagram) SaveAs(filename string) error {
	return saveAs(dia, filename)
}
