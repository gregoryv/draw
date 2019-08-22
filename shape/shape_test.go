package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_example_shapes(t *testing.T) {
	y := 70
	font := Font{Height: 9, Width: 7, LineHeight: 15}
	pad := Padding{Left: 10, Top: 2, Bottom: 7, Right: 10}

	addy := func(Y int) int {
		y += Y
		return y
	}
	cases := []struct {
		shape SvgWriter
	}{
		{
			&Svg{
				Width:  350,
				Height: 300,
				Content: []SvgWriterShape{
					&Line{X1: 0, Y1: y, X2: 100, Y2: y},
					&Arrow{X1: 230, Y1: y, X2: 280, Y2: y + 30}, // q1
					&Arrow{X1: 230, Y1: y, X2: 200, Y2: y + 50}, // q2
					&Arrow{X1: 230, Y1: y, X2: 180, Y2: y - 20}, // q3
					&Arrow{X1: 230, Y1: y, X2: 270, Y2: y - 20}, // q4
					&Arrow{X1: 230, Y1: y, X2: 320, Y2: y},      // right
					&Arrow{X1: 230, Y1: y, X2: 180, Y2: y},      // left
					&Arrow{X1: 230, Y1: y, X2: 230, Y2: y - 40}, // up
					&Arrow{X1: 230, Y1: y, X2: 230, Y2: y + 40}, // down

					&Label{
						Y:    addy(80),
						Text: "Label",
					},
					&Record{
						Y:     addy(20),
						Title: "Record",
						Font:  font,
						Pad:   pad,
						Fields: []string{
							"Write",
							"SetLabel",
						},
					},

					&Record{
						X:     100,
						Y:     y,
						Title: "Record without fields",
						Font:  font,
						Pad:   pad,
					},
				},
			},
		},
	}
	buf := bytes.NewBufferString("")
	for _, c := range cases {
		c.shape.WriteSvg(buf)
	}
	golden.Assert(t, buf.String())
}
