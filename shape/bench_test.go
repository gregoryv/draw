package shape

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkActor(b *testing.B) {
	shapes := []Shape{
		NewActor(),
		NewLabel("label"),
		NewRect("label"),
	}
	for _, shape := range shapes {
		b.Run(fmt.Sprintf("%T", shape), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				shape.WriteSVG(ioutil.Discard)
			}
		})
	}
}
