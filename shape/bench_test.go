package shape

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/draw"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &draw.Svg{}
	svg.Append(&Record{})

	style := NewStyle(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		svg.WriteSvg(&style)
	}
	b.StopTimer()
	b.ReportAllocs()
}
