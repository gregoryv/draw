package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

// NewCard returns a card with title, note and text.
func NewCard(title, note, text string) *Card {
	c := &Card{
		rect: NewRect(""),
	}
	c.SetClass("card")
	size := 18

	c.title = NewLabel(title)
	c.title.SetClass("card-title")
	c.title.Font.Height = size

	c.note = NewLabel(note)
	c.note.SetClass("card-note")

	c.text = NewLabel(text)
	c.text.Font.Height = 14

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

func (c *Card) SetTitle(v string) {
	c.title.Text = v
}

func (c *Card) SetNote(v string) {
	c.note.Text = v
}

func (c *Card) SetText(v string) {
	c.text.Text = v
}

func (c *Card) SetIcon(v Shape) {
	c.icon = v
}

func (c *Card) String() string {
	return fmt.Sprintf("Card %q", c.title.Text)
}
func (c *Card) Direction() Direction { return DirectionRight }

func (c *Card) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)

	T := c.title
	N := c.note
	D := c.text

	p := T.Font.Height
	c.rect.Pad.Top = p
	c.rect.Pad.Left = p
	c.rect.Pad.Bottom = p
	c.rect.Pad.Right = p

	c.rect.SetWidth(c.Width())
	c.rect.SetHeight(c.Height())
	c.rect.WriteSVG(w)

	top := c.rect.Pad.Top
	if c.icon != nil {
		NewAdjuster(c.icon).Below(c.rect, -c.Height()+c.rect.Pad.Top)
		new(Aligner).VAlignCenter(c.rect, c.icon)
		top += c.icon.Height()
		top += c.rect.Pad.Bottom
		c.icon.WriteSVG(w)
	}
	NewAdjuster(T).Below(c.rect, -c.Height()+top)
	NewAdjuster(N).Below(T, 0)
	NewAdjuster(D).Below(N, c.rect.Pad.Top)

	new(Aligner).VAlignCenter(c.rect, T, N, D)
	new(Aligner).VAlignLeft(c.rect, D)
	Move(D, c.rect.Pad.Left, 0)
	T.WriteSVG(w)
	N.WriteSVG(w)
	D.WriteSVG(w)

	return *err
}

func (c *Card) resize() {
	c.rect.SetWidth(c.Width())
	c.rect.SetHeight(c.Height())
}

func (c *Card) Width() int {
	if c.rect.width != 0 {
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
	if c.rect.height != 0 {
		return c.rect.Height()
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
