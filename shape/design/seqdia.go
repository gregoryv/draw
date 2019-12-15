package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/draw/shape"
)

// NewSequenceDiagram returns a sequence diagram with default column
// width.
func NewSequenceDiagram() *SequenceDiagram {
	return &SequenceDiagram{
		Diagram:  NewDiagram(),
		ColWidth: 190,
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
func (d *SequenceDiagram) WriteSvg(w io.Writer) error {
	var (
		colWidth = d.ColWidth

		top = d.top()
		x   = d.Pad.Left
		y1  = top + d.TextPad.Bottom + d.Font.LineHeight // below label
		y2  = d.Height()
	)
	lines := make([]*shape.Line, len(d.columns))
	for i, column := range d.columns {
		label := shape.NewLabel(column)
		label.Font = d.Font
		label.Pad = d.Pad
		label.SetX(i * colWidth)
		label.SetY(top)

		firstColumn := i == 0
		if firstColumn {
			x += label.Width() / 2
		}
		line := shape.NewLine(x, y1, x, y2)
		line.SetClass("column-line")
		lines[i] = line
		x += colWidth

		d.VAlignCenter(lines[i], label)
		d.Place(lines[i], label)
	}

	y := y1 + d.plainHeight()
	for _, lnk := range d.links {
		fromX := lines[lnk.fromIndex].Start.X
		toX := lines[lnk.toIndex].Start.X
		label := shape.NewLabel(lnk.text)
		label.Font = d.Font
		label.Pad = d.Pad
		label.SetX(fromX)
		label.SetY(y - 3 - d.Font.LineHeight)

		if lnk.toSelf() {
			margin := 15
			// add two lines + arrow
			l1 := shape.NewLine(fromX, y, fromX+margin, y)
			l1.SetClass(lnk.class())
			l2 := shape.NewLine(fromX+margin, y, fromX+margin, y+d.Font.LineHeight*2)
			l2.SetClass(lnk.class())
			d.HAlignCenter(l2, label)
			label.SetX(fromX + l1.Width() + d.TextPad.Left)
			label.SetY(y + 3)

			arrow := shape.NewArrow(
				l2.End.X,
				l2.End.Y,
				l1.Start.X,
				l2.End.Y,
			)
			arrow.SetClass(lnk.class())
			d.Place(l1, l2, arrow, label)
			y += d.selfHeight()
		} else {
			arrow := shape.NewArrow(
				fromX,
				y,
				toX,
				y,
			)
			arrow.SetClass(lnk.class())
			d.VAlignCenter(arrow, label)
			d.Place(arrow, label)
			y += d.plainHeight()
		}
	}
	return d.Diagram.WriteSvg(w)
}

// Width returns the total width of the diagram
func (d *SequenceDiagram) Width() int {
	if d.Svg.Width != 0 {
		return d.Svg.Width
	}
	return len(d.columns) * d.ColWidth
}

// Height returns the total height of the diagram
func (d *SequenceDiagram) Height() int {
	if d.Svg.Height != 0 {
		return d.Svg.Height
	}
	if len(d.columns) == 0 {
		return 0
	}
	height := d.top() + d.plainHeight()
	for _, lnk := range d.links {
		if lnk.toSelf() {
			height += d.selfHeight()
			continue
		}
		height += d.plainHeight()
	}
	return height
}

// selfHeight is the height of a self referencing link
func (d *SequenceDiagram) selfHeight() int {
	return 3*d.Font.LineHeight + d.Pad.Bottom
}

// plainHeight returns the height of and arrow and label
func (d *SequenceDiagram) plainHeight() int {
	return d.Font.LineHeight + d.Pad.Bottom + d.VMargin
}

func (d *SequenceDiagram) top() int {
	return d.Pad.Top
}

func (d *SequenceDiagram) AddColumns(names ...string) {
	d.columns = append(d.columns, names...)
}

func (d *SequenceDiagram) SaveAs(filename string) error {
	return saveAs(d, d.Style, filename)
}

func (d *SequenceDiagram) AddStruct(obj interface{}) string {
	name := reflect.TypeOf(obj).String()
	d.AddColumns(name)
	return name
}
