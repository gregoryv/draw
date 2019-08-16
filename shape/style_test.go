package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_parseClass(t *testing.T) {
	cases := []struct {
		line string
		exp  string
	}{
		{`class="hepp" x="8">`, "hepp"},
	}
	assert := asserter.New(t)

	for _, c := range cases {
		got := parseClass([]byte(c.line))
		assert().Equals(string(got), c.exp)
	}

}
