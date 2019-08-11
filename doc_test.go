package design

import (
	"testing"
)

func Test_render_design_document(t *testing.T) {
	doc := NewReferenceDoc()
	filename := "result.html"
	doc.SaveAs(filename)
}

func NewReferenceDoc() *DesignDoc {
	doc := NewDesignDoc()
	write := doc.Editor()
	write(
		"<h1>Go-Design Diagram Reference</h1>",
	)

	graph := NewGraph()
	graph.Title = "Struct component"
	graph.NewComponent(Account{})

	write(graph)
	return doc
}
