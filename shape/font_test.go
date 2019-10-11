package shape

import "testing"

func TestTextWidth(t *testing.T) {
	cases := []struct {
		r    rune
		font map[rune]float32
		exp  int
	}{
		{'ä', arial, 8},
	}
	for _, c := range cases {
		got := DefaultFont.TextWidth(string(c.r))
		if got != c.exp {
			t.Errorf("%q found=%v c.ok=%v", c.r, got, c.exp)
		}
	}

}
