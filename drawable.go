package design

import "io"

type Drawable interface {
	WriteTo(io.Writer) (int, error)
}

type Drawables []Drawable

func (all Drawables) WriteTo(w io.Writer) (int, error) {
	var total int
	for _, part := range all {
		n, err := part.WriteTo(w)
		if err != nil {
			return total + n, err
		}
		total += n
	}
	return total, nil
}
