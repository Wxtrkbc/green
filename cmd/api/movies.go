package main

import (
	"encoding/json"
	"fmt"
	"green/internal/data"
	"net/http"
	"time"
)

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
	//params := httprouter.ParamsFromContext(r.Context())
	//id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		//http.NotFound(w, r)
		app.notFoundResponse(w, r)
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
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}

}
