package shape

import "testing"

func TestOneSvg(t *testing.T) {
	it := &OneSvg{t, &Svg{}}
	// when
	it.IsEmpty()
	it.AppendsShapeAsFirstElementInContent()
	// after which
	it.AppendsShapesLastToContent()
	it.PrependsShapeFirstToContent()
}

type OneSvg struct {
	*testing.T
	*Svg
}

func (t *OneSvg) IsEmpty() {
	t.Helper()
	if len(t.Content) != 0 {
		t.Error("Not empty")
	}
}

func (t *OneSvg) AppendsShapeAsFirstElementInContent() {
	t.Helper()
	shape := &Arrow{}
	t.Append(shape)
	if t.Content[0] != shape {
		t.Error("Not first")
	}
}

func (t *OneSvg) AppendsShapesLastToContent() {
	t.Helper()
	shape := &Arrow{}
	t.Append(shape)
	if t.Content[len(t.Content)-1] != shape {
		t.Error("Not last")
	}
}

func (t *OneSvg) PrependsShapeFirstToContent() {
	t.Helper()
	shape := &Arrow{}
	t.Prepend(shape)
	if t.Content[0] != shape {
		t.Error("Not first")
	}
}
