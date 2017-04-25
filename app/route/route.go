package route

import (
	"net/http"

	"github.com/larsha/brynn.se-go/app/controller"
	"github.com/larsha/brynn.se-go/app/route/middleware/logrequest"

	"github.com/julienschmidt/httprouter"
)

func Load() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	r := httprouter.New()

	// Set 404 handler
	r.NotFound = http.HandlerFunc(controller.Error404)

	// Home page
	r.GET("/", controller.IndexGET)

	return r
}

func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Handler(h)

	return h
}
