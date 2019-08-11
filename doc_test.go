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
		"<h1>Example Go-Design Document</h1>",
		"<p>Some stuff here</p>",
	)

	account := NewComponent(Account{})
	ledger := NewComponent(Ledger{})
	product := NewComponent(Product{})

	graph := NewGraph()
	account.ShowFields()
	graph.Add(account)
	ledger.x = 200
	ledger.ShowFields()
	graph.Add(ledger)
	graph.Link(account, ledger)
	// graph.Link(comp, product) // invalid example

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

type Product struct{}
