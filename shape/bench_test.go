package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkWriteSvg(b *testing.B) {
	SvgWriter := &Svg{
		Content: []SvgWriterShape{
			&Record{},
		},
	}
	styler := &Styler{dest: ioutil.Discard}
	for i := 0; i < b.N; i++ {
		SvgWriter.WriteSvg(styler)
	}
	b.StopTimer()
	b.ReportAllocs()
}
