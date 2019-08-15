package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkWriteSvg(b *testing.B) {
	SvgWriter := &Svg{
		Content: []SvgWriter{
			&Record{},
		},
	}
	styler := &Styler{ioutil.Discard}
	for i := 0; i < b.N; i++ {
		SvgWriter.WriteSvg(styler)
	}
}
