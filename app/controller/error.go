package controller

import (
	"net/http"

	"github.com/larsha/brynn.se-go/app/shared/view"
)

// Error404 handles 404 - Page Not Found
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	v := view.New(r)
	v.Name = "error/404"
	v.Render(w)
}

// Error500 handles 500 - Internal Server Error
func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	v := view.New(r)
	v.Name = "error/500"
	v.Render(w)
}
