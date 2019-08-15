package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &Svg{
		Content: []svg{
			&Record{},
		},
	}
	styler := &Styler{ioutil.Discard}
	for i := 0; i < b.N; i++ {
		svg.WriteSvg(styler)
	}
}
