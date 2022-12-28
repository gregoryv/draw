package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

// NewCard returns a card with title, note and description.
func NewCard(title, note, description string) *Card {
	c := &Card{
		Rect: NewRect(""),
	}
	c.SetClass("card")
	size := 18

	c.Title = NewLabel(title)
	c.Title.SetClass("card-title")
	c.Title.Font.Height = size

	c.Note = NewLabel(note)
	c.Note.SetClass("card-note")

	c.Text = NewLabel(description)
	c.Text.Font.Height = 14

	return c
}

type Card struct {
	*Rect
	Title *Label
	Note  *Label
	Text  *Label
}

func (c *Card) String() string {
	return fmt.Sprintf("Card %q", c.Title.Text)
}
func (c *Card) Direction() Direction { return DirectionRight }

func (c *Card) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)

	T := c.Title
	N := c.Note
	D := c.Text

	p := T.Font.Height
	c.Rect.Pad.Top = p
	c.Rect.Pad.Left = p
	c.Rect.Pad.Bottom = p
	c.Rect.Pad.Right = p

	c.Rect.SetWidth(c.Width())
	c.Rect.SetHeight(c.Height())
	c.Rect.WriteSVG(w)

	NewAdjuster(T).Below(c.Rect, -c.Height()+c.Pad.Top)
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

func (c *Card) Width() int {
	if c.Rect.width != 0 {
		return c.Rect.Width()
	}
	width := c.Title.Width()
	if v := c.Note.Width(); v > width {
		width = v
	}
	if v := c.Text.Width(); v > width {
		width = v
		fmt.Println(v)
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
	h += c.Title.Height()
	h += c.Note.Height()
	h += c.Pad.Top
	h += c.Text.Height()
	h += c.Pad.Bottom

	return h
}
