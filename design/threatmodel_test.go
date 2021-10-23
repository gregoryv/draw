package design_test

import (
	"math"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/theme"
)

func Example_threatModel() {
	var tm CustomerProfilePage
	body := Body(
		Header(
			"Example of documenting a threat modelling in Go",
		),

		H1("Threat modelling"),
		P(),

		tm.Story(),
		tm.Model(),
		tm.Threats(),
		Hr(),
		Footer(
			`This example follows the article `,
			A(Href("https://martinfowler.com/articles/agile-threat-modelling.html"),
				"Agile Threat Modelling",
			),
			". ",
			`The goal is to find a suitable format for documenting the
            modelling session for future reuse.`,
		),
	)

	savePage(body)
	// output:
}

type CustomerProfilePage struct{}

func (me *CustomerProfilePage) Story() *Element {
	var (
		title, ref = "Customer profile page", ""
		content    = `As a customer, I need a page where I can see , my
			       customer details, So that I can confirm , they are
			       correct`
	)
	return Div(Class("story"),
		H2(title,
			Span(Class("ref"), ref),
		),
		P(content),
	)
}

func (me *CustomerProfilePage) Model() *Element {
	return Div(Class("model"),
		me.Diagram().Inline(),
	)
}

func (me *CustomerProfilePage) Diagram() *design.Diagram {
	var (
		d = design.NewDiagram()
		a = shape.NewActor()
		b = Entity("Customer\nDetails UI")
		e = Entity("Identity\nProvider")
		c = Entity("Customer\nDetails BFF")
		f = Entity("Customer\nService")
	)

	d.Style.Spacing = 60
	// Identify components
	d.Place(a).At(0, 140)
	d.Place(b, c).RightOf(a)
	d.Place(e).Above(c)
	d.Place(f).RightOf(c)

	// Add data flows
	d.LinkAll(a, b, c)
	d.Link(c, e)
	d.Link(c, f)

	// Identify trust boundaries
	b1 := Boundary(b, c, 20, -3)
	g := shape.NewLabel("internet")
	d.Place(b1)
	d.Place(g).At(40, 40)

	// Show your assets
	creds := Asset("creds")
	d.Place(creds).Above(e).Move(20, 80)

	pii := Asset("PII") // personally identifable information (PII)
	d.Place(pii).Above(f).Move(20, 80)

	return d
}

func (me *CustomerProfilePage) Threats() *Element {
	t := Table(
		Tr(
			Th("Data-flow"),
			Th("Threat"),
		),
		row(
			"Customer→Identity Service",
			"authentication is password based, no two-factor authentication",
		),
	)
	return t
}

// ----------------------------------------

func savePage(body *Element) {
	page := NewPage(
		Html(
			Head(
				Style(
					theme.GoldenSpace().With(
						theme.GoishColors(),
						modelTheme(),
					),
				),
			),
			body,
		),
	)
	page.SaveAs("showcase/threatmodel.html")
}

func row(flow string, threats ...string) *Element {
	return Tr(
		Td(flow),
		Td(
			func() *Element {
				ol := Ol()
				for _, threat := range threats {
					ol.With(Li(threat))
				}
				return ol
			}(),
		),
	)
}

func Asset(text string) shape.Shape {
	a := shape.NewRect(text)
	a.SetClass("asset")
	draw.DefaultClassAttributes["asset"] = `stroke="orange" fill="orange"`
	draw.DefaultClassAttributes["asset-title"] = `font-family="Arial,Helvetica,sans-serif"`
	return a
}

// Boundary
func Boundary(s1, s2 shape.Shape, extra, slant int) shape.Shape {
	x1, y1 := s1.Position()
	x2, y2 := s2.Position()

	xd := x2 - (x1 + s1.Width())
	x := x1 + s1.Width() + xd/2

	e := extra // extra
	s := slant // slant, a bit of an angle
	l := shape.NewLine(x+s, y1-e, x-s, y2+s2.Height()+e)
	l.SetClass("boundary")
	return l
}

func Entity(v string) shape.Shape {
	return shape.NewRect(v)
}

// Shapes follow e.g.
// https://docs.microsoft.com/en-us/learn/modules/tm-create-a-threat-model-using-foundational-data-flow-diagram-elements/1b-elements
//
// Workflow maps into
// https://martinfowler.com/articles/agile-threat-modelling.html
func intAbs(v int) int {
	return int(math.Abs(float64(v)))
}

func story(title string, ref string, lines ...interface{}) *Element {
	return Div(Class("story"),
		H2(title,
			Span(Class("ref"), ref),
		),
		P(lines...),
	)
}

func modelTheme() *CSS {
	css := NewCSS()
	css.Style(".story",
		"background-color: #f2f2f2",
		"padding: 5px 5px 5px 5px",
		"border-radius: 10px",
	)
	css.Style(".story h2",
		"font-size: 1em",
		"border-bottom: 1px solid #727272",
		"padding: 5px 5px 5px 5px",
		"margin-top: 0px",
	)
	css.Style(".story h2 .ref",
		"float: right",
		"font-size: 12px",
		"font-weight: normal",
	)
	css.Style(".story p",
		"padding: 5px 5px 5px 5px",
		"font-style: italic",
	)
	css.Style(".model",
		"padding: 1.612em 1.612em",
		"text-align: center",
	)
	return css
}

func figure(d *design.Diagram, caption string) *design.Diagram {
	d.SetCaption(caption)
	return d
}
