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
	collect := &ErrCollector{}
	collect.Last(fmt.Fprintf(w,
		`<svg width="%v" height="%v">`,
		shape.Width, shape.Height))

	for _, s := range shape.Content {
		fmt.Fprint(w, "\n")
		collect.Err(s.WriteSvg(w))
	}
	collect.Last(fmt.Fprint(w, "</svg>"))
	return collect.First()
}
