package design

import (
	"os"
	"testing"

	"github.com/gregoryv/go-design/shape"
)

func Test_example_sequence_diagram(t *testing.T) {
	diagram := &SequenceDiagram{
		Width:    500,
		Height:   200,
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
	diagram.Link(srv, srv, "Transform to view model")
	diagram.Link(srv, cli, "Send HTML")

	fh, err := os.Create("alldiagrams.svg")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	styler := shape.NewStyler(fh)
	diagram.WriteSvg(styler)
}
