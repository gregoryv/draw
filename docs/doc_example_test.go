package docs

import . "github.com/gregoryv/web"

func Example_RenderInlinedDiagram() {
	page := NewPage(
		Html(
			Body(
				H1("An example"),
				ExampleSmallClassDiagram().Inline(),
			),
		),
	)
	page.SaveAs("mypage.html")
}
