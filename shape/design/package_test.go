package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/shape"
)

func Test_saveAs(t *testing.T) {
	err := saveAs(&SequenceDiagram{}, shape.NewStyle(nil), "/")
	if err == nil {
		t.Fail()
	}
}

func checkInlining(t *testing.T, d draw.SVGInlineWriter) (inlined, plain bytes.Buffer) {
	t.Helper()
	d.InlineSVG(&inlined)
	d.WriteSVG(&plain)

	if inlined.String() == plain.String() {
		t.Error("Inlined same as plain")
		t.Log("inlined: \n", inlined.String())
		t.Log("plain: \n", plain.String())
	}
	return
}
