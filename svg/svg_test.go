package svg

import (
	"testing"
)

func Test_element_names(t *testing.T) {
	for e := Element_undefined; e <= Element_last+1; e++ {
		valid := e > Element_undefined && e < Element_last
		if valid && e.String() == "undefined" {
			t.Errorf("No names for %#v", e)
		}
	}
}
