package design

import (
	"os"
	"testing"
)

func ExampleSequenceDiagram() {
	diagram := NewSequenceDiagram()
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
	styler := NewStyler(fh)
	diagram.WriteSvg(styler)
	fh.Close()
}

func TestSequenceDiagram(t *testing.T) {
	ExampleSequenceDiagram()
}
