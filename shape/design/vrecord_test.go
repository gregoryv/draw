package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestVRecord(t *testing.T) {
	r := NewStruct(VRecord{})
	before := len(r.Fields)
	r.TitleOnly()
	got := len(r.Fields)
	assert := asserter.New(t)
	assert(got != before).Error("Did not hide fields")
}
