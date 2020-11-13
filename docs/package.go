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
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
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
	nav := Nav(H4("Table of contents"))
	article := Article(
		H1("Draw - programming software design diagrams"),

		P(`Draw Go module provides SVG rendering features focusing on
           creating software design diagrams.`),

		P(`Why? <br><em>Let programmers do what they are good
           at</em>. By providing an API, creating diagrams is no
           different from other programming. <br><em>Speed</em>; once you
           know the API programming a sequence diagram is actually
           much faster than drawing it by hand.  <br><em>Keep
           uptodate</em>; diagrams describe software, which moves
           fast. Keeping diagram in sync is a tedious task if manually
           drawn.  Being code, they are refactored with the same tools
           as the rest of the code and in most cases with little or no
           extra effort.`),
		Div(Class("left"),
			nav,
		),
		Div(Class("right"),
			H2("Install"),
			Pre(
				Code(
					"    go get ",
					A(Href("https://github.com/gregoryv/draw"), "github.com/gregoryv/draw/"),
					"...",
				),
			),

			H2("API documentation"),
			Ul(
				Li(
					A(Href(godoc("github.com/gregoryv/draw")), "draw"),
				),
				Li(
					A(Href(godoc("github.com/gregoryv/draw/shape")), "draw/shape"),
					" - SVG shapes",
				),
				Li(
					A(Href(godoc("github.com/gregoryv/draw/design")), "draw/design"),
					" - software design diagrams",
				),
			),
			H2("About"),
			Img(Src("me_circle.png"), Class("me")),
			P(
				`Written by `, A(Href("https://github.com/gregoryv"), gregory), Br(),
				A(Href("#license"), "MIT License"),
			),
		),
		Br(Attr("clear", "all")),

		H2("Example source"),
		Table(
			Tr(
				Td(
					LoadFile("small_example.go", 1, -1),
				),
				Td(Br(), Br(),
					ExampleSmallClassDiagram().Inline(),
				),
			),
		),

		H2("Diagrams"),
		H3("Class"),

		P(

			`In class diagrams the author wants to convey design
			relations between various entities. However the relations
			and most of the element naming can be generated from the
			source code. The author should add what is needed for a
			clear picture, ie. selecting entities to show and position
			them in a perceptible manner.`,
			//
		),
		ExampleClassDiagram().Inline(), Br(),
		"Source: ", A(Href("class_example.go"), "class_example.go"),
		P(

			`Records describe each entity using package name and
			type. Methods and fields are shown only by name if
			visible. Details such as arguments and return values are
			left to the API documentation. Relations are automatically
			rendererd between entities if there is one.`,
			//
		),

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

		H2("Changelog"),
		LoadFile("../changelog.md", 7, -1),

		H2("License"),
		LoadFile("../LICENSE"),
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
