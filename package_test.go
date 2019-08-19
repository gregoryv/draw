package design

import (
	"testing"
)

func Test_saveAs(t *testing.T) {
	err := saveAs(&SequenceDiagram{}, "/")
	if err == nil {
		t.Fail()
	}
}
