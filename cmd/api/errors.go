package main

import (
	"fmt"
	"net/http"
)

// generic log error
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// generic errResponse
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, 500, env, nil)
	if err != nil {
		app.logError(r, err)
	}

	// write to the header
	w.WriteHeader(500)
}

// server error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// a server error response needs to send an error message via the log
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// not found server error
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	// this error message is meant for the client and does not need to be logged
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// method not allowed error
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	// this error message is meant for the client and does not need to be logged
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
