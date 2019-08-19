package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestLink_class(t *testing.T) {
	cases := []struct {
		lnk Link
		exp string
	}{
		{Link{}, "arrow"},
		{Link{Class: "special"}, "special"},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := c.lnk.class()
		assert().Equals(got, c.exp)
	}
}

func TestSequenceDiagram_Link_missing_from_column(t *testing.T) {
	dia := &SequenceDiagram{}
	defer func() { recover() }()
	dia.Link("a", "b", "..")
	t.Fail()
}

func TestSequenceDiagram_Link_missing_to_column(t *testing.T) {
	dia := &SequenceDiagram{columns: []string{"a"}}
	defer func() { recover() }()
	dia.Link("a", "b", "..")
	t.Fail()
}
