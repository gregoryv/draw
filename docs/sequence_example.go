package docs

import (
	"database/sql"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/internal/app"
)

func ExampleSequenceDiagram() *design.SequenceDiagram {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})
		sqs = d.Add("aws.SQS")
	)
	d.Group(srv, sqs, "Private RPC using Gob encoding", "red")
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	d.Link(srv, srv, "Transform to view model").Class = "highlight"
	d.Link(srv, cli, "Send HTML")
	d.Link(srv, sqs, "Publish event")
	return d
}
