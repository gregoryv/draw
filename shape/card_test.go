package shape

import (
	"os"
	"testing"

	"github.com/gregoryv/draw"
)

func TestCard(t *testing.T) {
	customer := NewCard(
		"Personal Banking Customer",
		"[Person]",
		`A customer of the bank, with personal 
bank accounts.`,
	)
	customer.SetIcon(NewActor())
	customer.SetX(20)
	customer.SetY(20)
	customer.SetWidth(310)

	ibs := NewCard(
		// title
		"Internet Banking System",
		// type of thing
		"[Software System]",
		// description
		"Allows customers to view information about",
		"their bank accounts, and make payments.",
	)
	ibs.SetX(20)
	ibs.SetY(300)

	mailsys := NewCard(
		"E-mail System",
		"[Software System]",
		`The internal Microsoft Exchange
e-mail system.`,
	)
	mailsys.SetX(400)
	mailsys.SetY(300)
	mailsys.SetClass("card-external")

	d := &draw.SVG{}
	d.SetSize(800, 700)
	d.Append(ibs, mailsys, customer)

	filename := "testdata/card.svg"
	fh, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	style := draw.NewStyle()
	style.SetOutput(fh)
	d.WriteSVG(&style)
	fh.Close()

	testShape(t, ibs)
}
