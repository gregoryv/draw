package shape

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/golden"
)

func Test_tagPrinter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := newTagPrinter(buf)
	assert := asserter.New(t)
	assert(err != nil).Error(err)
	w.printf("ok %s\n", "printf")
	w.print("ok print\n")
	w.Write([]byte("ok Write\n"))

	w.err = fmt.Errorf("failed")
	w.printf("%s should not print this", "printf")
	w.print("print should not print this")
	w.Write([]byte("Write should not write this"))
	golden.Assert(t, buf.String())
}
