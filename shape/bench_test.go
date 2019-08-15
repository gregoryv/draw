package shape

import (
	"io/ioutil"
	"testing"
)

func BenchmarkWriteSvg(b *testing.B) {
	record := &Record{}
	styler := &Styler{ioutil.Discard}
	for i := 0; i < b.N; i++ {
		record.WriteSvg(styler)
	}
}
