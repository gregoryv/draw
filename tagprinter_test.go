package draw

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTagPrinter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := NewTagPrinter(buf)
	if err == nil {
		t.Fatal(err)
	}
	w.Printf("ok %s\n", "printf")
	w.Print("ok print\n")
	w.Write([]byte("ok Write\n"))

	w.err = fmt.Errorf("failed")
	w.Printf("%s should not print this", "printf")
	w.Print("print should not print this")
	w.Write([]byte("Write should not write this"))
	if strings.Index(buf.String(), " not ") > -1 {
		t.Error(buf.String())
	}
}
