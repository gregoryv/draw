package shape

import (
	"os"
	"testing"

	"github.com/gregoryv/draw"
)

func TestCard(t *testing.T) {
	ibs := NewCard(
		"Internet Banking System",
		"[Person]",
		`A customer of the bank, with personal
 bank accounts.`,
	)
	ibs.SetX(20)
	ibs.SetY(20)

	mailsys := NewCard(
		"E-mail System",
		"[Software System]",
		`The internal Microsoft Exchange
e-mail system.`,
	)
	mailsys.SetX(300)
	mailsys.SetY(200)
	mailsys.SetClass("card-external")

	d := &draw.SVG{}
	d.SetSize(800, 700)
	d.Append(ibs)
	d.Append(mailsys)

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
