package shape

import (
	"io"
)

type Record struct {
	X, Y         int
	Title        string
	PublicFields []string

	Font Font
	Pad  Padding
}

func (record *Record) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<rect class="record" x="%v" y="%v" width="%v" height="%v"/>`,
		record.X, record.Y, record.Width(), record.Height())
	w.printf("\n")
	record.writeFirstSeparator(w)
	var y = boxHeight(record.Font, record.Pad, 1) + record.Pad.Top
	for _, txt := range record.PublicFields {
		y += record.Font.LineHeight
		label := &Label{
			X:    record.X + record.Pad.Left,
			Y:    record.Y + y,
			Text: txt,
		}
		label.WriteSvg(w)
		w.printf("\n")
	}
	record.title().WriteSvg(w)
	return *err
}

func (record *Record) writeFirstSeparator(w io.Writer) error {
	if len(record.PublicFields) == 0 {
		return nil
	}
	y1 := record.Y + boxHeight(record.Font, record.Pad, 1)
	line := &Line{
		X1: record.X,
		Y1: y1,
		X2: record.X + record.Width(),
		Y2: y1,
	}
	return line.WriteSvg(w)
}

func (record *Record) title() *Label {
	return &Label{
		X:    record.X + record.Pad.Left,
		Y:    record.Y + record.Font.LineHeight + record.Pad.Top,
		Text: record.Title,
	}
}

func (record *Record) lines() int {
	return 1 + len(record.PublicFields)
}

func (record *Record) Height() int {
	first := boxHeight(record.Font, record.Pad, 1)
	if len(record.PublicFields) == 0 {
		return first
	}
	rest := boxHeight(record.Font, record.Pad, len(record.PublicFields))
	return first + rest
}

func (record *Record) Width() int {
	width := boxWidth(record.Font, record.Pad, record.Title)
	for _, txt := range record.PublicFields {
		w := boxWidth(record.Font, record.Pad, txt)
		if w > width {
			width = w
		}
	}
	return width
}

func (record *Record) Position() (int, int) { return record.X, record.Y }
func (record *Record) SetX(x int)           { record.X = x }
func (record *Record) SetY(y int)           { record.Y = y }
func (record *Record) Direction() Direction { return LR }
