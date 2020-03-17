package shape

import (
	"io"

	"github.com/gregoryv/draw"
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

func (d *Database) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	d.Cylinder.WriteSvg(w)

	l := &Label{
		x:     d.x + d.Pad.Left,
		y:     d.y + d.Pad.Top + int(d.ry()),
		Font:  d.Font,
		Text:  d.Title,
		class: d.class + "-title",
	}
	l.WriteSvg(w)

	return *err
}