package design

import (
	"testing"

	"github.com/gregoryv/draw/shape"
)

func Test_saveAs(t *testing.T) {
	err := saveAs(&SequenceDiagram{}, shape.NewStyle(nil), "/")
	if err == nil {
		t.Fail()
	}
}
