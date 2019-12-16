package design

import (
	"fmt"
	"io"
	"time"

	"github.com/gregoryv/draw/shape"
)

// NewGantChart returns a chart with number of days with optional
// start time. If none is given, then time.Now() is used.
func NewGanttChart(days int, start ...time.Time) *GanttChart {
	d := &GanttChart{
		Diagram: NewDiagram(),
		start:   time.Now(),
		days:    days,
		tasks:   make([]task, 0),
	}
	if len(start) > 0 {
		d.start = start[0]
	}
	return d
}

type GanttChart struct {
	Diagram
	start time.Time
	days  int
	tasks []task
}

func (d *GanttChart) Add(txt string, offset, days int) {
	d.tasks = append(d.tasks, task{txt, offset, days})
}

const indent int = 20 // top and left
const cw int = 16     // column width

type task struct {
	txt          string
	offset, days int
}

func (d *GanttChart) WriteSvg(w io.Writer) error {
	now := d.start
	d.Place(shape.NewLabel(fmt.Sprintf("%v", now.Year()))).At(indent, indent)
	offset := indent + d.daysOffset()
	for i := 0; i < d.days; i++ {
		x := offset + i*cw
		day := now.Day()
		if day == 1 {
			txt := fmt.Sprintf("%s", now.Month())
			d.Place(shape.NewLabel(txt)).At(x, indent)
		}
		col := shape.NewLabel(fmt.Sprintf("%02v", day))
		if now.Weekday() == time.Saturday {
			bg := shape.NewRect("")
			bg.SetClass("weekend")
			bg.SetWidth(cw*2 + 1)
			bg.SetHeight(len(d.tasks)*20 + 30)
			d.Place(bg).At(x+cw-3, 46)
		}

		col.Font.Height = 10
		d.Place(col).At(x, 40)
		now = now.AddDate(0, 0, 1)
	}

	for i, t := range d.tasks {
		y := i*20 + 70
		label := shape.NewLabel(t.txt)
		d.Place(label).At(indent, y)
		rect := shape.NewRect("")
		rect.SetWidth(t.days*cw - 5)
		rect.SetHeight(d.Diagram.Font.Height)
		rect.SetClass("span")
		d.Place(rect).At(offset+t.offset*cw, y)
		d.HAlignCenter(label, rect)
	}
	return d.Diagram.WriteSvg(w)
}

func (d *GanttChart) SaveAs(filename string) error {
	return saveAs(d, d.Diagram.Style, filename)
}

func (d *GanttChart) daysOffset() int {
	x := 0
	for _, t := range d.tasks {
		w := d.Diagram.Font.TextWidth(t.txt)
		if w > x {
			x = w
		}
	}
	return x + indent
}
