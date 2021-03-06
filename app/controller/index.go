package controller

import (
	"net/http"

	"github.com/larsha/fre.la/app/shared/view"

	"github.com/julienschmidt/httprouter"
)

type Index struct {
	Title string
}

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	v := view.New(r)
	v.Name = "index/page"
	v.Context = Index{Title: "Hello"}
	v.Render(w)
}
