package shape

import (
	"fmt"
	"io"
)

type Svg struct {
	Width, Height int
	Content       []svg
}

func (shape *Svg) WriteSvg(w io.Writer) error {
	e := make([]error, 2+len(shape.Content))
	_, e[0] = fmt.Fprintf(w,
		`<svg width="%v" height="%v">`,
		shape.Width, shape.Height)
	for i, s := range shape.Content {
		fmt.Fprint(w, "\n")
		e[i+1] = s.WriteSvg(w)
	}
	_, e[len(e)-1] = fmt.Fprint(w, "</svg>")
	return firstOf(e...)
}
