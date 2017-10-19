package controller

import (
	"net/http"

	"github.com/larsha/fre.la/app/shared/view"
)

// Error404 handles 404 - Page Not Found
func Error404(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "error/404"
	v.Status = http.StatusNotFound
	v.Render(w)
}

// Error500 handles 500 - Internal Server Error
func Error500(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "error/500"
	v.Status = http.StatusInternalServerError
	v.Render(w)
}
