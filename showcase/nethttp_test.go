package showcase

import (
	"net/http"
	"testing"

	design "github.com/gregoryv/go-design"
)

func TestClassDiagram_nethttp(t *testing.T) {
	var (
		d      = design.NewClassDiagram()
		h      = d.Interface((*http.Handler)(nil))
		r      = d.Struct(http.Request{})
		w      = d.Interface((*http.ResponseWriter)(nil))
		mux    = d.Struct(http.ServeMux{})
		server = d.Struct(http.Server{})
	)
	d.HideRealizations()

	d.Place(r).At(20, 20)
	d.Place(w).RightOf(r)
	d.Place(h).Below(w)
	d.Place(mux).Below(h)
	d.Place(server).Below(mux)

	d.SaveAs("nethttp.svg")
}
