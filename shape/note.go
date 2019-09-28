package shape

import (
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewNote(text string) *Note {
	return &Note{Text: text}
}

type Note struct {
	TopLeft xy.Position
	Text    string

	Font  Font
	Class string
}

func (note *Note) WriteSvg(out io.Writer) error {
	x1, y1 := note.TopLeft.XY()
	x2, y2 := note.TopLeft.XY() // todo use correct positions to draw a note like box
	w, err := newTagPrinter(out)
	w.printf(`<path class="%s" d="M%v,%v L%v,%v" />`, note.class(), x1, y1, x2, y2)
	return *err
}

// todo differentiate box class and text class, maybe label? make it multiline
func (note *Note) class() string {
	if note.Class == "" {
		return "note"
	}
	return note.Class
}

func (note *Note) Direction() Direction { return LR }
func (note *Note) Position() (int, int) { return note.TopLeft.XY() }
func (note *Note) SetX(x int)           { note.TopLeft.X = x }
func (note *Note) SetY(y int)           { note.TopLeft.Y = y }

func (note *Note) Height() int {
	return 50 // todo caculate height based on number of lines
}
func (note *Note) Width() int {
	return 100 // todo calculate width based on longest line
}
