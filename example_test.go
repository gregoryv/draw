package design

import (
	"os"
	"testing"

	"github.com/gregoryv/go-design/shape"
)

func ExampleSequenceDiagram() {
	diagram := &SequenceDiagram{
		Width:    500,
		Height:   230,
		ColWidth: 190,
		Font:     shape.Font{Height: 9, Width: 7, LineHeight: 15},
		TextPad:  shape.Padding{Left: 10, Top: 2, Bottom: 7, Right: 10},
		Pad:      shape.Padding{Left: 10, Top: 20, Bottom: 7, Right: 10},
	}
	cli, srv, db := "Client", "Server", "Database"
	diagram.AddColumns(cli, srv, db)
	diagram.Link(cli, srv, "connect()")
	diagram.Link(srv, db, "SELECT")
	diagram.Link(db, srv, "Rows")
	// Special link
	lnk := diagram.Link(srv, srv, "Transform to view model")
	lnk.Class = "highlight"
	diagram.Link(srv, cli, "Send HTML")

	// Save the diagram to file
	fh, _ := os.Create("img/sequence_example.svg")
	styler := shape.NewStyler(fh)
	diagram.WriteSvg(styler)
	fh.Close()
}

func TestSequenceDiagram(t *testing.T) {
	ExampleSequenceDiagram()
}
