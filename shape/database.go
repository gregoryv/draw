package shape

import (
	"io"

	"github.com/gregoryv/nexus"
)

func NewDatabase(title string) *Database {
	c := NewCylinder(40, 70)
	w := boxWidth(c.Font, c.Pad, title)
	c.height = boxHeight(c.Font, c.Pad, 2)
	c.Radius = w / 2

	c.SetClass("database")
	return &Database{
		Cylinder: c,
		Title:    title,
	}
}

type Database struct {
	*Cylinder
	Title string
}

func (d *Database) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	d.Cylinder.WriteSVG(w)

	l := &Label{
		x:     d.x + d.Pad.Left,
		y:     d.y + d.Pad.Top + int(d.ry()),
		Font:  d.Font,
		text:  d.Title,
		class: d.class + "-title",
	}
	l.WriteSVG(w)

	return *err
}
