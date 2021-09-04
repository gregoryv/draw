package design

import (
	"testing"

	"github.com/gregoryv/draw"
)

func Test_saveAs(t *testing.T) {
	err := saveAs(&SequenceDiagram{}, draw.NewStyle(), "/")
	if err == nil {
		t.Fail()
	}
}
