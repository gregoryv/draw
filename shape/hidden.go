package shape

import (
	"io"
)

// NewHidden returns a shap wrapper that is hidden, ie. not rendered.
func NewHidden(v Shape) *Hidden {
	return &Hidden{
		Shape: v,
	}
}

type Hidden struct {
	Shape
}

func (a *Hidden) WriteSVG(out io.Writer) error {
	return nil
}
