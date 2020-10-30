package docs

import (
	"bufio"
	"bytes"
	"os"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewProjectsPage() *Page {
	page := NewPage(Html(
		Head(
			Meta(Charset("utf-8")),
			Style(
				Theme(),
			),
		),
		Body(
			NewProjectArticle(),
		)))
	return page
}

func NewProjectArticle() *Element {
	nav := Nav("Table of contents")
	article := Article(
		H1("draw"),
		P(

			`Go module providing SVG rendering features. It comes with
             subpackages for shapes and software design diagrams. For
             software written in Go this module can help in keeping
             design diagrams uptodate. Diagrams is just Go-code, no
             special syntax is needed. Refactoring is applied together
             with your other code.  `,
			//
		),
		P(
			`Written by `, gregory,
		),
		nav,
		H2("API documentation"),
		Ul(
			Li(
				A(Href(godoc("github.com/gregoryv/draw")), "draw"),
			),
			Li(
				A(Href(godoc("github.com/gregoryv/draw/shape")), "draw/shape"),
			),
			Li(
				A(Href(godoc("github.com/gregoryv/draw/design")), "draw/design"),
			),
		),
		H2("Diagrams"),

		H3("Class"),
		ExampleClassDiagram().Inline(),
		"Source: ", A(Href("class_example.go"), "class_example.go"),

		H3("Activity"),
		ExampleActivityDiagram().Inline(),

		H3("Sequence"),
		ExampleSequenceDiagram().Inline(),

		H3("Gantt chart"),
		ExampleGanttChart().Inline(),

		H3("Generic"),
		ExampleDiagram().Inline(),

		H2("Shapes"),
		AllShapes().Inline(),
	)
	toc.GenerateIDs(article, "h2", "h3")
	nav.With(toc.ParseTOC(article, "h2", "h3"))
	return article
}

func godoc(pkg string) string {
	return "https://godoc.org/" + pkg
}

const gregory = "Gregory Vin&ccaron;i&cacute;"

func AllShapes() *design.Diagram {
	d := design.NewDiagram()
	vspace := 60

	actorLbl := shape.NewLabel("Actor")
	actor := shape.NewActor()

	d.Place(actorLbl).At(20, 20)
	d.Place(actor).RightOf(actorLbl, vspace+40)

	lastLabel := actorLbl
	var last shape.Shape = actor
	add := func(txt string, s shape.Shape) {
		label := shape.NewLabel(txt)
		d.Place(label, s).Below(lastLabel, vspace)
		d.VAlignCenter(last, s)
		d.HAlignCenter(label, s)
		lastLabel = label
		last = s
	}

	add("Arrow", shape.NewArrow(240, 0, 300, 0))
	add("Circle", shape.NewCircle(20))
	add("Component", shape.NewComponent("Component"))
	add("Cylinder", shape.NewCylinder(30, 40))
	add("Database", shape.NewDatabase("database"))
	add("Diamond", shape.NewDiamond())
	add("Dot", shape.NewDot())
	add("ExitDot", shape.NewExitDot())
	add("Internet", shape.NewInternet())
	add("Label", shape.NewLabel("label-text"))
	shape.Move(last, 0, -18) // todo labels do not align properly
	add("Line", shape.NewLine(240, 0, 300, 0))
	add("Note", shape.NewNote("This describes\nsomething..."))

	rec := shape.NewRecord("record")
	rec.Fields = []string{"fields"}
	rec.Methods = []string{"methods"}
	add("Record", rec)

	add("Rect", shape.NewRect("a rectangle"))
	add("State", shape.NewState("active"))
	add("Triangle", shape.NewTriangle())

	return &d
}

// LoadFile returns a pre web element wrapping the contents from the
// given file. If to == -1 all lines to the end of file are returned.
func LoadFile(filename string, span ...int) *Element {
	from, to := 0, -1
	if len(span) == 2 {
		from, to = span[0], span[1]
	}
	v := loadFile(filename, from, to)
	class := "srcfile"
	if from == 0 && to == -1 {
		class += " complete"
	}
	return Pre(Class(class), Code(Class("go"), v))
}

func loadFile(filename string, from, to int) string {
	var buf bytes.Buffer
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fh)
	for i := from; i > 1; i-- {
		scanner.Scan()
		to--
	}

	for scanner.Scan() {
		to--
		buf.WriteString(scanner.Text() + "\n")
		if to == 0 {
			break
		}
	}
	return buf.String()
}
