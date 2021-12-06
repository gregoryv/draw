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
	*Diagram
	ColWidth int
	VMargin  int // top margin for each horizontal lane

	columns []string
	links   []*Link
	groups  []group
}

type group struct {
	fromColumn string
	toColumn   string
	text       string
	class      string
}

// Group adds a colored area below span of columns. Predefined classes
// are red, green, blue.
func (me *SequenceDiagram) Group(fromColumn, toColumn, text, class string) {
	g := group{
		fromColumn: fromColumn,
		toColumn:   toColumn,
		text:       text,
		class:      class,
	}
	if me.groups == nil {
		me.groups = []group{g}
		return
	}
	me.groups = append(me.groups, g)
}

// WriteSvg renders the diagram as SVG to the given writer.
func (d *SequenceDiagram) WriteSVG(w io.Writer) error {
	var (
		colWidth = d.ColWidth

		top = d.top()
		x   = d.Pad.Left
		y1  = top + d.TextPad.Bottom + d.Font.LineHeight // below label
		y2  = d.Height()
	)
	lines := make([]*shape.Line, len(d.columns))
	vlines := make(map[string]*shape.Line)
	// save x values for rendering skip lines
	columnX := make([]int, len(d.columns))
	// columns and vertical lines
	for i, column := range d.columns {
		// todo add label in Add, AddStruct and AddInterface methods
		// so one can place other shapes relative to them
		label := shape.NewLabel(column)
		label.Font = d.Font
		label.Pad = d.Pad
		label.SetX(i * colWidth)
		label.SetY(top)

		firstColumn := i == 0
		if firstColumn {
			x += label.Width() / 2
			columnX = append(columnX, x)
		}
		line := shape.NewLine(x, y1, x, y2)
		line.SetClass("column-line")
		lines[i] = line
		x += colWidth
		columnX = append(columnX, x)

		d.VAlignCenter(lines[i], label)
		d.Place(lines[i], label)
		vlines[column] = line // save for groups

		// groups
		for _, group := range d.groups {
			if group.toColumn == column { // assume from is already there
				r := shape.NewRect("") // add align label

				alabel := shape.NewLabel(group.text)
				alabel.Font = d.Font
				alabel.Pad = d.Pad
				alabel.SetClass("area-" + group.class + "-label")

				from := vlines[group.fromColumn]
				to := line
				x, y := from.Position()
				x2, _ := to.Position()
				width := x2 - x
				r.SetX(x)
				r.SetY(y)
				r.SetWidth(width)
				r.SetClass("area-" + group.class)
				r.SetHeight(y2 + d.top() + label.Height())
				d.Prepend(r) // behind

				d.Place(alabel).Below(r)
				d.VAlignCenter(r, alabel)
				d.HAlignBottom(r, alabel)
				shape.Move(alabel, 0, -alabel.Pad.Bottom)
			}
		}
	}

	y := y1 + d.plainHeight()
	for _, lnk := range d.links {
		if lnk == skip {
			for _, x := range columnX {
				dots := shape.NewLine(x, y, x, y+d.Font.LineHeight)
				dots.SetClass("skip")
				d.Place(dots)
			}
			y += d.plainHeight()
			continue
		}
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
	return d.Diagram.WriteSVG(w)
}

// Width returns the total width of the diagram
func (d *SequenceDiagram) Width() int {
	w := d.SVG.Width()
	if w != 0 {
		return w
	}
	return len(d.columns) * d.ColWidth
}

// Height returns the total height of the diagram
func (d *SequenceDiagram) Height() int {
	h := d.SVG.Height()
	if h != 0 {
		return h
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

// AddColumns adds the names as columns in the given order.
func (d *SequenceDiagram) AddColumns(names ...string) {
	for _, name := range names {
		d.Add(name)
	}
}

// Add the name as next column and return name.
func (d *SequenceDiagram) Add(name string) string {
	d.columns = append(d.columns, name)
	return name
}

func (d *SequenceDiagram) SaveAs(filename string) error {
	return saveAs(d, d.Style, filename)
}

// Inline returns rendered SVG with inlined style
func (d *SequenceDiagram) Inline() string {
	return inline(d, d.Style)
}

// String returns rendered SVG
func (d *SequenceDiagram) String() string { return toString(d) }

func (d *SequenceDiagram) AddStruct(obj interface{}) string {
	name := reflect.TypeOf(obj).String()
	d.Add(name)
	return name
}

func (d *SequenceDiagram) AddInterface(obj interface{}) string {
	name := reflect.TypeOf(obj).Elem().String()
	d.Add(name)
	return name
}
