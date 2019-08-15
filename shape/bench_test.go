package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkSvgRendering(b *testing.B) {
	record := &Record{}
	for i := 0; i < b.N; i++ {
		record.Svg()
	}
}

func BenchmarkWriteSvg(b *testing.B) {
	record := &Record{}
	for i := 0; i < b.N; i++ {
		record.WriteSvg(ioutil.Discard)
	}
}
