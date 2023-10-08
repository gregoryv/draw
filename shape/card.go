package shape

import (
	"io"
	"strings"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

// NewCard returns a card with title optional note and text. Default
// width is set to 310px. You can use multiple lines for text which
// will be joined with newlines.
func NewCard(title string, args ...string) *Card {
	size := 18
	p := size
	r := NewRect("")
	r.SetWidth(310)
	r.Pad.Top = p
	r.Pad.Left = p
	r.Pad.Bottom = p
	r.Pad.Right = p

	c := &Card{
		rect: r,
	}
	c.SetClass("card")

	t := NewLabel(title)
	t.SetClass("card-title")
	t.Font.Height = size
	t.Pad.Left = 0
	t.Pad.Right = 0
	c.title = t

	if len(args) > 0 {
		note := args[0]
		c.note = NewLabel(note)
		c.note.SetClass("card-note")
	} else {
		c.note = NewLabel("")
	}

	if len(args) > 1 {
		text := strings.Join(args[1:], "\n")
		c.text = NewLabel(text)
		c.text.Font.Height = 14
	} else {
		c.text = NewLabel("")
	}

	return c
}

type Card struct {
	rect  *Rect
	title *Label
	note  *Label
	text  *Label

	// Optional icon placed above the title
	icon Shape
}

func (c *Card) SetClass(v string)    { c.rect.SetClass(v) }
func (c *Card) SetX(v int)           { c.rect.SetX(v) }
func (c *Card) SetY(v int)           { c.rect.SetY(v) }
func (c *Card) SetWidth(v int)       { c.rect.SetWidth(v) }
func (c *Card) SetHeight(v int)      { c.rect.SetHeight(v) }
func (c *Card) Position() (x, y int) { return c.rect.Position() }

func (c *Card) Edge(start xy.Point) xy.Point { return boxEdge(start, c) }

func (c *Card) SetTitle(v string) { c.title.SetText(v) }
func (c *Card) SetNote(v string)  { c.note.SetText(v) }
func (c *Card) SetText(v string)  { c.text.SetText(v) }
func (c *Card) SetIcon(v Shape)   { c.icon = v }

func (c *Card) Direction() Direction { return DirectionRight }

func (c *Card) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)

	T := c.title
	N := c.note
	D := c.text

	c.rect.width = c.Width()
	c.rect.height = c.Height()
	c.rect.WriteSVG(w)

	halfpad := c.rect.Pad.Left / 2
	top := c.rect.Pad.Top
	if c.icon != nil {
		NewAdjuster(c.icon).Below(c.rect, -c.Height()+c.rect.Pad.Top)
		new(Aligner).VAlignCenter(c.rect, c.icon)
		top += c.icon.Height()
		top += c.rect.Pad.Bottom
		Move(c.icon, -halfpad, 0)
		c.icon.WriteSVG(w)
	}
	NewAdjuster(T).Below(c.rect, -c.Height()+top)
	NewAdjuster(N).Below(T, 0)
	NewAdjuster(D).Below(N, c.rect.Pad.Top)

	new(Aligner).VAlignCenter(c.rect, T, N, D)
	new(Aligner).VAlignLeft(c.rect, D)
	Move(D, c.rect.Pad.Left, 0)

	// todo figure out why VAlignCenter doesn't quite work here

	Move(T, -halfpad, 0)
	Move(N, -halfpad, 0)

	T.WriteSVG(w)
	N.WriteSVG(w)
	D.WriteSVG(w)

	return *err
}

func (c *Card) Width() int {
	if c.rect.width > 0 {
		return c.rect.Width()
	}
	width := c.title.Width()
	if v := c.note.Width(); v > width {
		width = v
	}
	if v := c.text.Width(); v > width {
		width = v
	}
	width += c.rect.Pad.Left
	width += c.rect.Pad.Right
	return width
}

func (c *Card) Height() int {
	if c.rect.height > 0 {
		return c.rect.height
	}
	h := c.rect.Pad.Top
	if c.icon != nil {
		h += c.icon.Height()
		h += c.rect.Pad.Bottom
	}
	h += c.title.Height()
	h += c.note.Height()
	h += c.rect.Pad.Top
	h += c.text.Height()
	h += c.rect.Pad.Bottom
	return h
}
