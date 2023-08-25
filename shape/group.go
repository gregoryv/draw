package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
)

// NewGroup returns a virtual group of shapes which can be moved
// together.
func NewGroup(shapes ...Shape) *Group {
	return &Group{
		Shapes: shapes,
		Pad:    Padding{Top: 12, Left: 12, Bottom: 12, Right: 12},
	}
}

type Group struct {
	Shapes []Shape
	Pad    Padding
}

func (g *Group) String() string {
	return fmt.Sprintf("Group")
}

func (g *Group) Position() (x, y int) {
	return g.TopLeftPos()
}

// SetX moves all shapes withing a group.
func (g *Group) SetX(x int) {
	cx, _ := g.TopLeftPos()
	for _, s := range g.Shapes {
		Move(s, x-cx, 0)
	}
}

// SetY moves all shapes withing a group.
func (g *Group) SetY(y int) {
	_, cy := g.TopLeftPos()
	for _, s := range g.Shapes {
		Move(s, 0, y-cy)
	}
}

func (g *Group) Direction() Direction { return DirectionRight }

// SetClass is a noop
func (g *Group) SetClass(c string) {}

// WriteSVG is a noop
func (g *Group) WriteSVG(out io.Writer) error { return nil }

func (g *Group) SetPad(pad Padding) { g.Pad = pad }

func (g *Group) Height() int {
	_, minY := g.TopLeftPos()
	_, maxY := g.BottomRightPos()
	return maxY - minY
}

func (g *Group) Width() int {
	minX, _ := g.TopLeftPos()
	maxX, _ := g.BottomRightPos()
	return maxX - minX
}

func (g *Group) TopLeftPos() (x, y int) {
	for _, s := range g.Shapes {
		sx, sy := s.Position()
		if x == 0 || sx < x {
			x = sx
		}
		if y == 0 || sy < y {
			y = sy
		}
	}
	x -= g.Pad.Left
	y -= g.Pad.Top

	return
}

func (g *Group) BottomRightPos() (x, y int) {
	for _, s := range g.Shapes {
		sx, sy := s.Position()
		w, h := s.Width(), s.Height()
		if x == 0 || sx+w > x {
			x = sx + w
		}
		if y == 0 || sy+h > y {
			y = sy + h
		}
	}
	x += g.Pad.Right
	y += g.Pad.Bottom

	return
}

// Edge returns intersecting position of a line starting at start and
// pointing to the rect center.
func (g *Group) Edge(start xy.Point) xy.Point {
	return boxEdge(start, g)
}
