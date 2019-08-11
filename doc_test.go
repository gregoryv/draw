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
	comp.ShowFields()
	graph.Add(comp)

	ledger := NewComponent(Ledger{})
	ledger.x = 200
	ledger.ShowFields()

	graph.Add(ledger)

	write(graph)
	return doc
}

type Account struct {
	Username string
	password string
}

type Ledger struct {
	From, To *Account
	Total    int
}
