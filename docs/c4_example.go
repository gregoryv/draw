package docs

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleC4SystemContextDiagram() *design.Diagram {
	var (
		d = design.NewDiagram()

		person = shape.NewCard(
			"Customer", "[Person]",
			"Someone buying books",
		)
		store  = shape.NewCard(
			"Book Store",
			"[Software System]",
			"Register of all books.",
		)
		acc    = shape.NewCard(
			"Accounting",
			"[Software System]",
			"Journal of all sales",
		)

		cont = shape.NewContainer(
			shape.NewLabel("[System Context] Book store"),
			person, store, acc,
		)
	)
	d.Spacing = 150
	person.SetIcon(shape.NewActor())
	acc.SetClass("external")

	d.Place(cont).At(10, 10)
	d.Place(person).At(20, 20)
	d.Place(acc).Below(person)
	d.Place(store).RightOf(acc)

	d.Link(person, store, "Finds and buys a book")
	d.Link(store, acc, "Logs purchases")
	return d
}
