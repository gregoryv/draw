package design

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_render_design_document(t *testing.T) {
	doc := NewDesignDoc()
	write := doc.Editor()
	write("<h1>title</h1>")

	filename := "result.html"
	doc.SaveAs(filename)
	golden.Assert(t, loadFile(t, filename))
}

func loadFile(t *testing.T, filename string) string {
	t.Helper()
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(body)
}
