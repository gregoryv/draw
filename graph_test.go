package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_components_diagram(t *testing.T) {
	graph := NewGraph()
	graph.Title = "Example of struct representation"
	graph.NewComponent(Account{})
	buf := bytes.NewBufferString("")
	graph.WriteTo(buf)
	golden.Assert(t, buf.String())
}

type Account struct {
	Username string
	password string
}
