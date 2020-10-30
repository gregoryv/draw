package docs

import "testing"

func Test_render_pages(t *testing.T) {
	index := NewProjectsPage()
	index.SaveAs("index.html")
}
