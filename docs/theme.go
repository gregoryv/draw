package docs

import . "github.com/gregoryv/web"

func Theme() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style("article",
		"margin: 1em 1.62em",
	)
	css.Style("*",
		"font-family: sans-serif",
	)
	css.Style("h1,h2,h3,h4,h5,h6",
		"font-family: serif",
	)
	css.Style("p",
		"line-height: 1.5em",
	)
	css.Style("nav>ul",
		"list-style-type: none",
		"line-height: 1.5em",
	)
	css.Style("li.h3",
		"margin-left: 1em",
		"list-style-type: none",
	)
	return css
}
