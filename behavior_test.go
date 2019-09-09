package design

import "testing"

func Test_one_sequence_diagram(t *testing.T) {
	it := new_one_sequence_diagram(t)
	it.is_empty()
	// when
	it.has_columns()
	it.has_no_links()
	it.is_not_empty()
	// and if
	it.has_fixed_size()
	it.is_rendered_with_fixed_size()
}
