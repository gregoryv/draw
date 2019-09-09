package design

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/go-design/shape"
	"github.com/gregoryv/go-design/style"
)

func BenchmarkWriteSvg(b *testing.B) {
	diagram := &shape.Svg{
		Content: []shape.SvgWriterShape{
			&shape.Record{},
		},
	}
	styler := style.NewStyler(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		diagram.WriteSvg(styler)
	}
	b.StopTimer()
	b.ReportAllocs()
}
