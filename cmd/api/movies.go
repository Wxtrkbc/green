package main

import (
	"fmt"
	"green/internal/data"
	"net/http"
	"time"
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

	var movie = data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = app.writeJsonResponse(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
