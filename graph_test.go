package design_test

import (
	"testing"

	"github.com/gregoryv/asserter"
	design "github.com/gregoryv/go-design"
)

func Test_linked(t *testing.T) {
	account := design.NewComponent(Account{})
	ledger := design.NewComponent(Ledger{})
	product := design.NewComponent(Product{})

	cases := []struct {
		a, b *design.Component
		exp  bool
	}{
		{account, ledger, true},
		{account, product, false},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := c.a.AreLinked(c.b)
		assert().Equals(got, c.exp)
	}
}
