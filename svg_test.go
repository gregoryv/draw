package design

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_element_names(t *testing.T) {
	buf := bytes.NewBufferString("")
	var e Element
	for e = Element_undefined; e <= Element_last+1; e++ {
		line := fmt.Sprintf("%s\n", e)
		buf.WriteString(line)
	}
	golden.Assert(t, buf.String())
}
