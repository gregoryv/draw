package design

import (
	"io/ioutil"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_render_design_document(t *testing.T) {
	doc := NewDesignDoc()
	write := doc.Editor()
	write("<h1>Design Document as Software</h1>")

	graph := NewGraph()
	graph.Title = "Struct component"
	graph.NewComponent(Account{})

	write(graph)

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
