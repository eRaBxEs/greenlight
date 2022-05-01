package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) exampleJsonEncoderHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hello": "world"}
	// Set the "Content-Type: application/json" header on the response.
	w.Header().Set("Content-Type", "application/json")
	// Use the json.NewEncoder() function to initialize a json.Encoder instance that
	// writes to the http.ResponseWriter. Then we call its Encode() method, passing in // the data that we want to encode to JSON (which in this case is the map above). If // the data can be successfully encoded to JSON, it will then be written to our
	// http.ResponseWriter.
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
