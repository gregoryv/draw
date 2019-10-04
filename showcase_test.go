package design_test

import (
	"net/http"
	"testing"

	design "github.com/gregoryv/go-design"
	"github.com/gregoryv/go-design/shape"
)

func TestClassDiagram_nethttp(t *testing.T) {
	var (
		d     = design.NewClassDiagram()
		title = shape.NewLabel("package net/http")
		h     = d.Interface((*http.Handler)(nil))

		client    = d.Struct(http.Client{})
		cookie    = d.Struct(http.Cookie{})
		cjar      = d.Interface((*http.CookieJar)(nil))
		req       = d.Struct(http.Request{})
		resp      = d.Struct(http.Response{})
		w         = d.Interface((*http.ResponseWriter)(nil))
		mux       = d.Struct(http.ServeMux{})
		server    = d.Struct(http.Server{})
		transport = d.Struct(http.Transport{})
	)
	d.HideRealizations()

	d.Place(title).At(220, 20)
	d.Place(h).At(20, 60)
	d.Place(mux).RightOf(h)
	d.Place(w).RightOf(mux)
	d.Place(req).RightOf(w)

	d.Place(server).Below(h)
	d.Place(resp).Below(mux)
	d.HAlignBottom(resp, server)

	d.Place(transport).Below(w)
	d.Place(client).Below(transport)
	d.Place(cookie).Below(server)
	d.Place(cjar).RightOf(cookie)

	d.SaveAs("img/http_classdiagram.svg")
}
