package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
	//params := httprouter.ParamsFromContext(r.Context())
	//id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)

}
