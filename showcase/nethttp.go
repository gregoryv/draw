package showcase

import (
	"net/http"

	design "github.com/gregoryv/go-design"
)

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

type App struct{}

type Index struct{}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
