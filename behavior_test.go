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

func Test_one_arrow(t *testing.T) {
	it := new_one_arrow(t)
	// when
	it.starts_at_visible_position_N()
	it.ends_above_and_right_of_N()
	it.points_up_and_right()
}
