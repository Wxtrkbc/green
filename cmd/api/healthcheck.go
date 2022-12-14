package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status": "available", "environment": %q, "version": %q}`
	//js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJsonResponse(w, http.StatusOK, data, nil)
	//js, err := json.Marshal(data)
	if err != nil {
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}
	//
	//w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body.
	//w.Write([]byte(js))
	//w.Write(js)

}
