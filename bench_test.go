package draw_test

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/shape"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &draw.SVG{}
	svg.Append(&shape.Record{})

	style := draw.NewStyle(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		svg.WriteSVG(&style)
	}
	b.StopTimer()
	b.ReportAllocs()
}
