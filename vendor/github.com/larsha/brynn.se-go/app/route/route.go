package route

import (
	"net/http"

	"github.com/larsha/brynn.se-go/app/controller"
	"github.com/larsha/brynn.se-go/app/route/middleware/logrequest"
	"github.com/larsha/brynn.se-go/app/shared/config"

	"github.com/julienschmidt/httprouter"
)

func Load() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	config := config.Get()
	r := httprouter.New()

	if config.Production == false {
		// Static
		r.ServeFiles(config.Static+"/*filepath", http.Dir("static"))
	}

	// Set 404 handler
	r.NotFound = http.HandlerFunc(controller.Error404)

	// Home page
	r.GET("/", controller.IndexGET)

	// Api
	r.POST("/api/form", controller.FormPOST)

	return r
}

func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Handler(h)

	return h
}
