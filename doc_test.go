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

	var (
		account = NewComponent(Account{})
		ledger  = NewComponent(Ledger{})
		product = NewComponent(Product{})
		order   = NewComponent(Order{})
	)

	graph := NewGraph()
	graph.Place(account.WithFields(), 0, 0)
	graph.Place(ledger.WithFields(), 200, 0)
	graph.Place(product, 200, 200)
	graph.Place(order.WithFields(), 0, 200)

	graph.Link(account, ledger)
	graph.Link(order, product)
	graph.Link(order, account)
	//graph.Link(account, product) // invalid example

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

type Order struct {
	First *Product
	Owner *Account
}
