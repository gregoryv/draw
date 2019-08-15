package shape

import (
	"fmt"
	"io"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

func (line *Line) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<line x1="%v" y1="%v" x2="%v" y2="%v"/>`,
		line.X1, line.Y1, line.X2, line.Y2,
	)
	return err
}
