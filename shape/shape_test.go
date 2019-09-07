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
					// q1
					NewArrow(230, y, 280, y+30),
					// q2
					NewArrow(230, y, 200, y+50),
					// q3
					NewArrow(230, y, 180, y-20),
					// q4
					NewArrow(230, y, 270, y-20),
					// right
					NewArrow(230, y, 320, y),
					// left
					NewArrow(230, y, 180, y),
					// up
					NewArrow(230, y, 230, y-40),
					//down
					NewArrow(230, y, 230, y+40),

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
