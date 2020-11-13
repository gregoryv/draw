package docs

import "io/ioutil"

var w = ioutil.Discard

type House struct {
	Frontdoor Door      // aggregation
	Windows   []*Window // composition
}

func (me *House) Rooms() int { return 0 }

type Door struct {
	Material string
}

func (me *Door) Materials() []string {
	return []string{}
}

type Window struct {
	Model string
}

func (me *Window) Materials() []string {
	return []string{}
}

type Part interface {
	Materials() []string
}
