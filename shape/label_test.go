package shape

import (
	"bytes"
	"testing"
)

func TestLabel_SetHref(t *testing.T) {
	var (
		a      = NewLabel("hepp")
		before bytes.Buffer
		after  bytes.Buffer
		_      = a.WriteSVG(&before)
	)
	a.SetHref("https://gregoryv.github.io/draw")
	a.WriteSVG(&after)
	if before.String() == after.String() {
		t.Error("SetHref had no effect")
		t.Log("before:", before.String())
	}
}
