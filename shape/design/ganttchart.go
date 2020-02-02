package design

import (
	"fmt"
	"io"
	"time"

	"github.com/gregoryv/draw/internal/date"
	"github.com/gregoryv/draw/shape"
)

// NewGanttChart returns a GanttChart spanning days from the given
// date. Panics if date cannot be resolved.
func NewGanttChart(from date.String, days int) *GanttChart {
	return newGanttChart(from.Time(), days)
}

// newGanttChart returns a chart showing days from optional
// start time. If no start is given, time.Now() is used.
func newGanttChart(start time.Time, days int) *GanttChart {
	d := &GanttChart{
		Diagram:  NewDiagram(),
		start:    start,
		days:     days,
		tasks:    make([]*Task, 0),
		padLeft:  16,
		padTop:   10,
		colSpace: 4,
		Mark:     time.Now(),
	}
	return d
}

type GanttChart struct {
	Diagram
	start time.Time
	days  int
	tasks []*Task

	padLeft, padTop int
	colSpace        int // between day or week

	// Set a marker at this date.
	Mark time.Time
}

func (d *GanttChart) MarkDate(yyyymmdd date.String) {
	d.Mark = yyyymmdd.Time()
}

// isToday returns true if time.Now matches start + ndays
func (d *GanttChart) isToday(ndays int) bool {
	t := d.start.AddDate(0, 0, ndays)
	return t.Year() == d.Mark.Year() &&
		t.YearDay() == d.Mark.YearDay()
}

// Add new task from start spanning 3 days. Default color is green.
func (d *GanttChart) Add(txt string) *Task {
	task := NewTask(txt, 0, 3)
	d.tasks = append(d.tasks, task)
	return task
}

// Add new task. Default color is green.
func (d *GanttChart) Place(task *Task) *GanttAdjuster {
	return &GanttAdjuster{
		start: d.start,
		task:  task,
	}
}

type GanttAdjuster struct {
	start time.Time
	task  *Task
}

func (a *GanttAdjuster) At(from date.String, days int) {
	a.task.offset = from.DaysAfter(a.start)
	if a.task.offset < 0 {
		a.task.offset = 0
	}
	a.task.days = days
}

func (a *GanttAdjuster) After(parent *Task, days int) {
	a.task.offset = parent.offset + parent.days
	a.task.days = days
}

func (d *GanttChart) WriteSvg(w io.Writer) error {
	columns := d.addHeader()
	bars := make([]*shape.Rect, len(d.tasks))
	start := d.padLeft + d.taskWidth()
	lineHeight := d.Diagram.Font.LineHeight
	headerHeight := d.padTop + lineHeight*3
	for i, t := range d.tasks {
		rect := shape.NewRect("")
		rect.SetHeight(d.Diagram.Font.Height)
		rect.SetClass(t.class)
		bars[i] = rect
		d.drawTask(i, t)
		y := i*lineHeight + headerHeight
		d.Diagram.Place(rect).At(start, y)
	}

	for i := 0; i < d.days; i++ {
		dayWidth := columns[i].Width() + d.colSpace
		for j, t := range d.tasks {
			switch {
			case i == t.offset:
				bars[j].SetX(start + i*dayWidth)
			case i > t.offset && i < t.offset+t.days:
				bars[j].SetWidth(bars[j].Width() + dayWidth)
			}
		}
	}

	return d.Diagram.WriteSvg(w)
}

func (d *GanttChart) addHeader() []*shape.Label {
	now := d.start
	year := shape.NewLabel(fmt.Sprintf("%v", now.Year()))
	d.Diagram.Place(year).At(d.padLeft, d.padTop)
	offset := d.padLeft + d.taskWidth()

	var lastDay *shape.Label
	columns := make([]*shape.Label, d.days)
	for i := 0; i < d.days; i++ {
		day := now.Day()
		col := newCol(day)
		columns[i] = col
		if now.Weekday() == time.Saturday {
			bg := shape.NewRect("")
			bg.SetClass("weekend")
			bg.SetWidth((col.Width() + d.colSpace) * 2)
			bg.SetHeight(len(d.tasks)*col.Font.LineHeight + d.padTop + d.Diagram.Font.LineHeight + d.colSpace)
			d.Diagram.Place(bg).RightOf(lastDay, d.colSpace)
			shape.Move(bg, -d.colSpace/2, d.colSpace)
		}
		if i == 0 {
			d.Diagram.Place(col).Below(year, d.colSpace)
			col.SetX(offset)
		} else {
			d.Diagram.Place(col).RightOf(lastDay, d.colSpace)
		}
		if day == 1 {
			label := shape.NewLabel(now.Month().String())
			d.Diagram.Place(label).Above(col, d.colSpace)
		}
		if d.isToday(i) {
			x, y := col.Position()
			mark := shape.NewLine(x, y, x+10, y)
			d.Diagram.Place(mark)
		}
		lastDay = col
		now = now.AddDate(0, 0, 1)
	}
	return columns
}

func (d *GanttChart) drawTask(i int, t *Task) {
	label := shape.NewLabel(t.txt)
	lineHeight := d.Diagram.Font.LineHeight
	headerHeight := d.padTop + lineHeight*3
	x := d.padLeft
	y := i*lineHeight + headerHeight - lineHeight/3
	d.Diagram.Place(label).At(x, y)
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
func (t *Task) Red() *Task { t.class = "span-red"; return t }

// Blue sets class of task to span-blue
func (t *Task) Blue() *Task { t.class = "span-blue"; return t }
