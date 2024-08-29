package design_test

import (
	"testing"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func TestC4Example(t *testing.T) {
	var (
		d        = design.NewDiagram()
		customer = shape.NewCard("Personal Banking Customer", "[Person]",
			`A customer of the bank, with personal 
bank accounts.`,
		)
		ibs = shape.NewCard("Internet Banking System", "[Software System]",
			`Allows customers to view information
about their bank accounts, and make
payments.`,
		)
		mailsys = shape.NewCard("E-mail System", "[Software System]",
			`The internal Microsoft Exchange
e-mail system.`,
		)
		mainframe = shape.NewCard("Mainframe Banking System", "[Software System]",
			`Stores all of the core banking
information about customers,
accounts, transactions, etc.`,
		)
	)

	mailsys.SetClass("card-external")
	mainframe.SetClass("card-external")
	customer.SetIcon(shape.NewActor())

	d.Place(customer).At(20, 20)
	d.Place(ibs).Below(customer, 170)
	d.Place(mailsys).RightOf(ibs, 200)
	d.Place(mainframe).Below(ibs, 170)
	d.HAlignCenter(ibs, mailsys)
	d.VAlignCenter(customer, ibs, mainframe)

	d.Link(customer, ibs,
		"Views account\nbalances, and\nmakes payments\nusing",
	)
	d.Link(ibs, mailsys, "Sends e-mail\nusing")
	d.Link(mailsys, customer, "Sends e-mail to")
	d.Link(ibs, mainframe,
		"Gets account\ninformation from,\nand makes\npayments using",
	)
	d.SetCaption("C4 example diagram")

	if err := d.SaveAs("img/c4_example.svg"); err != nil {
		t.Fatal(err)
	}
}
