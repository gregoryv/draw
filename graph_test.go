package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_linked(t *testing.T) {
	account := NewComponent(Account{})
	ledger := NewComponent(Ledger{})
	product := NewComponent(Product{})

	cases := []struct {
		a, b *Component
		exp  bool
	}{
		{account, ledger, true},
		{account, product, false},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := c.a.areLinked(c.b)
		assert().Equals(got, c.exp)
	}
}
