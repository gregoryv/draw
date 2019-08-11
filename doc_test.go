package design

import (
	"testing"
)

func Test_reference_document(t *testing.T) {
	NewReferenceDoc().SaveAs("reference.html")
}

func NewReferenceDoc() *DesignDoc {
	var (
		account = NewComponent(Account{})
		ledger  = NewComponent(Ledger{})
		product = NewComponent(Product{})
		order   = NewComponent(Order{})
	)

	graph := NewGraph()
	graph.Place(account.WithFields()).At(1, 40)
	graph.Place(ledger.WithFields()).RightOf(account)
	graph.Place(order.WithFields()).Below(account)
	graph.Place(product).RightOf(order)

	AlignHorizontal(Center, account, ledger)
	AlignHorizontal(Center, order, product)
	AlignVertical(Center, account, order)
	AlignVertical(Center, ledger, product)

	graph.Link(account, ledger)
	graph.Link(order, product)
	graph.Link(order, account)
	//graph.Link(account, product) // invalid example

	doc := NewDesignDoc()
	write := doc.Editor()
	write(
		"<h1>Example Go-Design Document</h1>",
		"<p>", Account{}, "s are the root connector of ", Order{}, "'s</p>",
		graph,
	)
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
