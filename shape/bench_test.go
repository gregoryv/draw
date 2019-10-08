package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &Svg{
		Content: []Shape{
			&Record{},
		},
	}
	style := NewStyle(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		svg.WriteSvg(&style)
	}
	b.StopTimer()
	b.ReportAllocs()
}
