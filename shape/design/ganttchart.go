package design

import (
	"fmt"
	"io"
	"time"

	"github.com/gregoryv/draw/shape"
)

func NewGanttChartFrom(days, yyyy, mm, dd int) *GanttChart {
	str := fmt.Sprintf("%v-%v-%vT01:00:00.000Z", yyyy, mm, dd)
	t, _ := time.Parse(time.RFC3339, str)
	return NewGanttChart(days, t)
}

// NewGanttChart returns a chart showing days from optional
// start time. If no start is given, time.Now() is used.
func NewGanttChart(days int, start ...time.Time) *GanttChart {
	d := &GanttChart{
		Diagram: NewDiagram(),
		start:   time.Now(),
		days:    days,
		tasks:   make([]*Task, 0),
		padLeft: 16,
		padTop:  10,
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
	tasks []*Task

	padLeft, padTop int
}

// Add new task. Default color is green.
func (d *GanttChart) Add(txt string, offset, days int) *Task {
	task := NewTask(txt, offset, days)
	d.tasks = append(d.tasks, task)
	return task
}

// NewTask returns a green task.
func NewTask(txt string, offset, days int) *Task {
	return &Task{
		txt:    txt,
		offset: offset,
		days:   days,
		class:  "span-green",
	}
}

// Task is the colorized span of a gantt chart.
type Task struct {
	txt          string
	offset, days int
	class        string
}

// Red sets class of task to span-red
func (t *Task) Red() { t.class = "span-red" }

// Blue sets class of task to span-blue
func (t *Task) Blue() { t.class = "span-blue" }

func (d *GanttChart) WriteSvg(w io.Writer) error {
	now := d.start
	year := shape.NewLabel(fmt.Sprintf("%v", now.Year()))
	d.Place(year).At(d.padLeft, d.padTop)
	offset := d.padLeft + d.taskWidth()

	var lastDay *shape.Label
	columns := make([]*shape.Label, d.days)
	for i := 0; i < d.days; i++ {
		day := now.Day()
		if day == 1 {
			label := shape.NewLabel(now.Month().String())
			d.Place(label).Above(lastDay, 4)
			shape.Move(label, lastDay.Width()+4, 0)
		}
		col := newCol(day)
		columns[i] = col
		if now.Weekday() == time.Saturday {
			bg := shape.NewRect("")
			bg.SetClass("weekend")
			bg.SetWidth(col.Width()*2 + 8)
			bg.SetHeight(len(d.tasks)*col.Font.LineHeight + d.padTop + d.Diagram.Font.LineHeight)
			d.Place(bg).RightOf(lastDay, 4)
			shape.Move(bg, -2, 4)
		}
		if i == 0 {
			d.Place(col).Below(year, 4)
			col.SetX(offset)
		} else {
			d.Place(col).RightOf(lastDay, 4)
		}
		lastDay = col
		now = now.AddDate(0, 0, 1)
	}

	var lastTask *shape.Label
	for i, t := range d.tasks {
		label := shape.NewLabel(t.txt)
		if i == 0 {
			d.Place(label).Below(lastDay, 4)
			d.VAlignLeft(year, label)
		} else {
			d.Place(label).Below(lastTask, 4)
		}
		lastTask = label

		rect := shape.NewRect("")
		col := columns[t.offset]
		var w int
		for i := t.offset; i < t.offset+t.days; i++ {
			w += columns[i].Width() + 4
		}
		rect.SetWidth(w - 4)
		rect.SetHeight(d.Diagram.Font.Height)
		rect.SetClass(t.class)

		d.Place(rect).Below(col, 4)
		d.HAlignCenter(label, rect)
	}
	return d.Diagram.WriteSvg(w)
}

func newCol(day int) *shape.Label {
	col := shape.NewLabel(fmt.Sprintf("%02v", day))
	col.Font.Height = 10
	return col
}

func (d *GanttChart) SaveAs(filename string) error {
	return saveAs(d, d.Diagram.Style, filename)
}

func (d *GanttChart) taskWidth() int {
	x := 0
	for _, t := range d.tasks {
		w := d.Diagram.Font.TextWidth(t.txt)
		if w > x {
			x = w
		}
	}
	return x + d.padLeft
}
