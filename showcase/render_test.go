package showcase

import (
	"fmt"
	"net/http"
	"testing"

	design "github.com/gregoryv/go-design"
)

func TestRenderAllDiagrams(t *testing.T) {
	cases := []struct {
		d        saveCap
		filename string
		caption  string
	}{
		{BasicNetHttpClassDiagram(), "nethttp.svg", "ServeMux routes requests to handlers"},
		{BackendHandler(), "backend.svg", "ServeMux is the router"},
		{BackendClassDiagram(), "backend_classes.svg", "Index implements http.Handler"},
	}
	for i, c := range cases {
		c.d.SetCaption(fmt.Sprintf("Figure %v. %s", i+1, c.caption))
		c.d.SaveAs(c.filename)
	}
}

type saveCap interface {
	SetCaption(string)
	SaveAs(string) error
}

func BasicNetHttpClassDiagram() *design.ClassDiagram {
	var (
		d      = design.NewClassDiagram()
		h      = d.Interface((*http.Handler)(nil))
		r      = d.Struct(http.Request{})
		w      = d.Interface((*http.ResponseWriter)(nil))
		mux    = d.Struct(http.ServeMux{})
		server = d.Struct(http.Server{})
	)
	d.HideRealizations()
	r.TitleOnly()

	d.Place(r).At(20, 20)
	d.Place(w).RightOf(r)
	d.Place(h).Below(w)
	d.Place(mux).Below(h)
	d.VAlignCenter(w, h, mux)

	d.Place(server).RightOf(w)
	return d
}

func BackendHandler() *design.SequenceDiagram {
	var (
		d   = design.NewSequenceDiagram()
		app = d.AddStruct(App{})
		h   = d.AddStruct(Index{})
		mux = d.AddStruct(http.ServeMux{})
		srv = d.AddStruct(http.Server{})
		cli = d.AddStruct(http.Client{})
	)
	d.ColWidth = 130
	d.Link(app, h, `&Index{} : myhandler`)
	d.Link(app, mux, `Handle("/path", myhandler)`)
	d.Link(app, srv, `ListenAndServe(":8080", mux)`)
	d.Link(cli, srv, `GET /path `)
	d.Link(srv, mux, `routes request to registered func`)
	return d
}

func BackendClassDiagram() *design.ClassDiagram {
	var (
		d = design.NewClassDiagram()
		h = d.Interface((*http.Handler)(nil))
		i = d.Struct(Index{})
	)
	d.HideRealizations()

	d.Place(h).At(20, 20)
	d.Place(i).Below(h)
	return d
}

type App struct{}

type Index struct{}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
