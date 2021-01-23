package design

import (
	"fmt"
	"io"
	"time"

	"github.com/gregoryv/draw"
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
	*Diagram
	start time.Time
	days  int
	tasks []*Task

	padLeft, padTop int
	colSpace        int // between day or week

	// Set a marker at this date.
	Mark time.Time

	Weeks bool
}

func (g *GanttChart) MarkDate(yyyymmdd date.String) {
	g.Mark = yyyymmdd.Time()
}

// isToday returns true if time.Now matches start + ndays
func (g *GanttChart) isToday(ndays int) bool {
	t := g.start.AddDate(0, 0, ndays)
	sameYear := t.Year() == g.Mark.Year()
	sameDay := t.YearDay() == g.Mark.YearDay()
	return sameYear && sameDay
}

// Add new task from start spanning 3 days. Default color is green.
func (g *GanttChart) Add(txt string) *Task {
	task := NewTask(txt)
	g.tasks = append(g.tasks, task)
	return task
}

// Add new task. Default color is green.
func (g *GanttChart) Place(task *Task) *GanttAdjuster {
	return &GanttAdjuster{
		start: g.start,
		task:  task,
	}
}

type GanttAdjuster struct {
	start time.Time
	task  *Task
}

func (a *GanttAdjuster) At(from date.String, days int) {
	a.task.from = from.Time()
	a.task.to = from.Time().AddDate(0, 0, days)
}

func (a *GanttAdjuster) After(parent *Task, days int) {
	a.task.from = parent.to
	a.task.to = parent.to.AddDate(0, 0, days)
}

func (d *GanttChart) WriteSVG(w io.Writer) error {
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
	// adjust the bars

	for j, t := range d.tasks {
		var width int
		for i := 0; i < d.days; i++ {
			now := d.start.AddDate(0, 0, i)
			var col *shape.Label
			if d.Weeks {
				col = columns[i/7]
			} else {
				col = columns[i]
			}
			switch {
			case now.Equal(t.from):
				d.VAlignLeft(col, bars[j])
			case now.After(t.from) && now.Before(t.to) || now.Equal(t.to):
				if !d.Weeks || (d.Weeks && now.Weekday() == time.Sunday) {
					width += col.Width()
					width += d.colSpace
				}
			}
		}
		if !d.Weeks {
			width -= d.colSpace
		}
		bars[j].SetWidth(width)
	}

	return d.Diagram.WriteSVG(w)
}

func (d *GanttChart) addHeader() []*shape.Label {
	now := d.start
	year := shape.NewLabel(fmt.Sprintf("%v", now.Year()))
	d.Diagram.Place(year).At(d.padLeft, d.padTop)
	offset := d.padLeft + d.taskWidth()

	var lastDay *shape.Label
	columns := make([]*shape.Label, 0)
	var col *shape.Label
	for i := 0; i < d.days; i++ {
		day := now.Day()
		colName := day
		if d.Weeks {
			_, colName = now.ISOWeek()
		}
		wday := now.Weekday()
		if day == 1 || d.Weeks && wday == 1 || !d.Weeks {
			col = newCol(colName)
			columns = append(columns, col)

			if !d.Weeks && now.Weekday() == time.Saturday {
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
				monthName := now.Month().String()
				if d.Weeks {
					monthName = monthName[:3]
				}
				label := shape.NewLabel(monthName)
				d.Diagram.Place(label).Above(col, d.colSpace)
			}
			lastDay = col
		}
		if d.isToday(i) {
			x, y := col.Position()
			mark := shape.NewLine(x, y, x+10, y)
			d.Diagram.Place(mark)
		}
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

// Inline returns rendered SVG with inlined style
func (d *GanttChart) Inline() string {
	return draw.Inline(d, d.Diagram.Style)
}

// String returns rendered SVG
func (d *GanttChart) String() string { return toString(d) }

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
func NewTask(txt string) *Task {
	return &Task{
		txt:   txt,
		class: "span-green",
	}
}

// Task is the colorized span of a gantt chart.
type Task struct {
	txt      string
	from, to time.Time
	class    string
}

// Red sets class of task to span-red
func (t *Task) Red() *Task { t.class = "span-red"; return t }

// Blue sets class of task to span-blue
func (t *Task) Blue() *Task { t.class = "span-blue"; return t }
