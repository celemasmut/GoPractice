package controller

import (
	"encoding/json"
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//once we run the host we'll see code a Body in json
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(data)
		}
	}
}
