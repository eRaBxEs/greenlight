package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) exampleJsonEncoderHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hello": "world",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and so could not process your request", http.StatusInternalServerError)
	}

}
