package design

import (
	"testing"
)

func Test_reference_document(t *testing.T) {
	NewReferenceDoc().SaveAs("reference.html")
}

func NewReferenceDoc() *DesignDoc {
	doc := NewDesignDoc()
	write := doc.Editor()
	write(
		"<h1>Go-Design Diagram Reference</h1>",
	)

	graph := NewGraph()
	graph.Title = "Struct component"
	comp := NewComponent(Account{})
	comp.ShowPublicFields()
	graph.Add(comp)

	write(graph)
	return doc
}

type Account struct {
	Username string
	password string
}
