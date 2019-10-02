package design

import (
	"io"

	"github.com/gregoryv/go-design/shape"
)

// NewSequenceDiagram returns a sequence diagram with default column
// width.
func NewSequenceDiagram() *SequenceDiagram {
	return &SequenceDiagram{
		ColWidth: 190,
		Diagram:  NewDiagram(),
		VMargin:  10,
	}
}

// SequenceDiagram defines columns and links between columns.
type SequenceDiagram struct {
	Diagram
	ColWidth int
	VMargin  int // top margin for each horizontal lane

	columns []string
	links   []*Link
}

// WriteSvg renders the diagram as SVG to the given writer.
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
		label := shape.NewLabel(column)
		label.X = i * colWidth
		label.Y = top
		label.Font = dia.Font
		label.Pad = dia.Pad

		firstColumn := i == 0
		if firstColumn {
			x += label.Width() / 2
		}
		line := shape.NewLine(x, y1, x, y2)
		line.SetClass("column-line")
		lines[i] = line
		x += colWidth

		dia.VAlignCenter(lines[i], label)
		svg.Content = append(svg.Content, lines[i], label)
	}

	y := y1 + dia.plainHeight()
	for _, lnk := range dia.links {
		fromX := lines[lnk.fromIndex].Start.X
		toX := lines[lnk.toIndex].Start.X
		label := shape.NewLabel(lnk.text)
		label.X = fromX
		label.Y = y - 2
		label.Font = dia.Font
		label.Pad = dia.Pad

		if lnk.toSelf() {
			margin := 15
			// add two lines + arrow
			l1 := shape.NewLine(fromX, y, fromX+margin, y)
			l1.SetClass(lnk.class())
			l2 := shape.NewLine(fromX+margin, y, fromX+margin, y+dia.Font.LineHeight*2)
			l2.SetClass(lnk.class())
			dia.HAlignCenter(l2, label)
			label.X += l1.Width() + dia.TextPad.Left
			arrow := shape.NewArrow(
				l2.End.X,
				l2.End.Y,
				l1.Start.X,
				l2.End.Y,
			)
			arrow.SetClass(lnk.class())
			svg.Content = append(svg.Content, l1, l2, arrow, label)
			y += dia.selfHeight()
		} else {
			arrow := shape.NewArrow(
				fromX,
				y,
				toX,
				y,
			)
			arrow.SetClass(lnk.class())
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
	return dia.Font.LineHeight + dia.Pad.Bottom + dia.VMargin
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
