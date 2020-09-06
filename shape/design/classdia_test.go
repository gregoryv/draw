package design

import (
	"testing"

	"github.com/gregoryv/draw/shape"
)

func TestClassDiagram(t *testing.T) {
	t.Run("InlineSVG", func(t *testing.T) {
		d := NewClassDiagram()
		record := d.Struct(shape.Record{})
		d.Place(record).At(20, 20)
		checkInlining(t, d)
	})
}
