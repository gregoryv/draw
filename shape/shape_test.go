package shape

import (
	"bytes"
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_shapes_write_svg(t *testing.T) {
	cases := []struct {
		shape     StyledWriter
		xmlText   string
		styleText string
	}{
		{&Line{}, "<line", "stroke"},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		buf := bytes.NewBufferString("")
		c.shape.WriteTo(buf)
		assert().Contains(buf.String(), c.xmlText)
		assert().Contains(c.shape.Style(), c.styleText)
	}
}

type StyledWriter interface {
	io.WriterTo
	Style() string
}
