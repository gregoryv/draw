package shape

import (
	"os"
	"testing"

	"github.com/gregoryv/draw"
)

func TestNote(t *testing.T) {
	n := NewNote(`Multiline text is
possible in notes`)
	n.SetY(20)
	saveAsSvg(t, n, "testdata/note.svg")
	testShape(t, n)

}

func saveAsSvg(t *testing.T, shape Shape, filename string) {
	t.Helper()
	d := &draw.SVG{}
	d.SetSize(300, 100)
	d.Append(shape)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	style := NewStyle(fh)
	d.WriteSvg(&style)
	fh.Close()
}
