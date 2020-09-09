package design_test

import (
	"database/sql"
	"testing"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/internal/app"
)

func ExampleSequenceDiagram() {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})
	)
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	d.Link(srv, srv, "Transform to view model").Class = "highlight"
	d.Link(srv, cli, "Send HTML")
	d.SaveAs("img/app_sequence_diagram.svg")
}

func TestExample(t *testing.T) {
	ExampleSequenceDiagram()
}
