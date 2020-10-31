package docs

import . "github.com/gregoryv/web"

func Theme() *CSS {
	css := NewCSS()
	css.Import("https://fonts.googleapis.com/css2?family=Inconsolata&family=Source+Sans+Pro&family=Source+Serif+Pro:wght@600")
	css.Import("https://fonts.googleapis.com/css2?family=Inconsolata&family=Lancelot&family=Open+Sans")
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
		"font-family: 'Source Sans Pro', sans-serif",
	)
	css.Style("article",
		"margin: 1em 1.62em",
	)
	css.Style("h1, h2, h3, h4, h5, h6",
		"font-family: 'Source Serif Pro', serif",
	)
	css.Style(".writtenby",
		"float: right",
	)
	css.Style(".toc",
		"font-weight: bold",
	)
	css.Style("p, li",
		"font-family: 'Open Sans', sans-serif",
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
	css.Style("td",
		"vertical-align: top",
	)

	// source code
	css.Style(".srcfile",
		"padding: .6em 1.6em .6em 1.6em",
		"display: block",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-color: #eaeaea",
	)

	css.Style("code",
		"font-family: Inconsolata",
		"-moz-tab-size: 4",
	)

	css.Style(".left",
		"float: left",
		"margin-right: 6em",
	)
	css.Style(".right",
		"float: left",
	)
	return css
}
