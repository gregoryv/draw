package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

// NewCard returns a card with title, note and text.
func NewCard(title, note, text string) *Card {
	c := &Card{
		Rect: NewRect(""),
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
	*Rect
	title *Label
	note  *Label
	text  *Label

	// Optional icon placed above the title
	icon Shape
}

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
	c.Rect.Pad.Top = p
	c.Rect.Pad.Left = p
	c.Rect.Pad.Bottom = p
	c.Rect.Pad.Right = p

	c.Rect.SetWidth(c.Width())
	c.Rect.SetHeight(c.Height())
	c.Rect.WriteSVG(w)

	top := c.Pad.Top
	if c.icon != nil {
		NewAdjuster(c.icon).Below(c.Rect, -c.Height()+c.Pad.Top)
		new(Aligner).VAlignCenter(c.Rect, c.icon)
		top += c.icon.Height()
		top += c.Pad.Bottom
		c.icon.WriteSVG(w)
	}
	NewAdjuster(T).Below(c.Rect, -c.Height()+top)
	NewAdjuster(N).Below(T, 0)
	NewAdjuster(D).Below(N, c.Pad.Top)

	new(Aligner).VAlignCenter(c.Rect, T, N, D)
	new(Aligner).VAlignLeft(c.Rect, D)
	Move(D, c.Pad.Left, 0)
	T.WriteSVG(w)
	N.WriteSVG(w)
	D.WriteSVG(w)

	return *err
}

func (c *Card) resize() {
	c.Rect.SetWidth(c.Width())
	c.Rect.SetHeight(c.Height())
}

func (c *Card) Width() int {
	if c.Rect.width != 0 {
		return c.Rect.Width()
	}
	width := c.title.Width()
	if v := c.note.Width(); v > width {
		width = v
	}
	if v := c.text.Width(); v > width {
		width = v
	}
	width += c.Rect.Pad.Left
	width += c.Rect.Pad.Right
	return width
}

func (c *Card) Height() int {
	if c.Rect.height != 0 {
		return c.Rect.Height()
	}
	h := c.Pad.Top
	if c.icon != nil {
		h += c.icon.Height()
		h += c.Pad.Bottom
	}
	h += c.title.Height()
	h += c.note.Height()
	h += c.Pad.Top
	h += c.text.Height()
	h += c.Pad.Bottom

	return h
}
