package draw

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/golden"
)

func Test_tagPrinter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := NewTagPrinter(buf)
	assert := asserter.New(t)
	assert(err != nil).Error(err)
	w.Printf("ok %s\n", "printf")
	w.Print("ok print\n")
	w.Write([]byte("ok Write\n"))

	w.err = fmt.Errorf("failed")
	w.Printf("%s should not print this", "printf")
	w.Print("print should not print this")
	w.Write([]byte("Write should not write this"))
	golden.Assert(t, buf.String())
}
