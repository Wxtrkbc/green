package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routers() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMoviesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies:id", app.showMoviesHandler)

	return router
}
