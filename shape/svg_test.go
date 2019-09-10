package shape

import "testing"

func TestSvg(t *testing.T) {
	it := &one_svg{t, &Svg{}}
	// when
	it.is_empty()
	it.appends_shape_as_first_element_in_content()
	// after which
	it.appends_shapes_last_to_content()
	it.prepends_shape_first_to_content()
}

type one_svg struct {
	*testing.T
	*Svg
}

func (t *one_svg) is_empty() {
	t.Helper()
	if len(t.Content) != 0 {
		t.Error("Not empty")
	}
}

func (t *one_svg) appends_shape_as_first_element_in_content() {
	t.Helper()
	shape := &Arrow{}
	t.Append(shape)
	if t.Content[0] != shape {
		t.Error("Not first")
	}
}

func (t *one_svg) appends_shapes_last_to_content() {
	t.Helper()
	shape := &Arrow{}
	t.Append(shape)
	if t.Content[len(t.Content)-1] != shape {
		t.Error("Not last")
	}
}

func (t *one_svg) prepends_shape_first_to_content() {
	t.Helper()
	shape := &Arrow{}
	t.Prepend(shape)
	if t.Content[0] != shape {
		t.Error("Not first")
	}
}
