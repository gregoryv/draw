package shape

import (
	"os"
	"testing"

	"github.com/gregoryv/go-design/style"
)

func TestNote(t *testing.T) {
	note := NewNote(`Multiline text
is possible in notes`)
	saveAsSvg(t, note, "testdata/note.svg")

}

func saveAsSvg(t *testing.T, shape SvgWriterShape, filename string) {
	t.Helper()
	d := &Svg{Width: 100, Height: 100}
	d.Append(shape)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(style.NewStyler(fh))
	fh.Close()
}
