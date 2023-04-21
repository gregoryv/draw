package goviz_test

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gregoryv/draw/goviz"
)

func ExampleSaveFrameSeq() {
	h := func(w http.ResponseWriter, r *http.Request) {
		goviz.SaveFrameSeq("testdata/gorilla.svg")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", h)
	go http.ListenAndServe(":9991", r)
	<-time.After(10 * time.Millisecond)
	http.Get("http://localhost:9991")
}
