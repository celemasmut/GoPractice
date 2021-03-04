package main

import (
	"encoding/json"
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}

//once we run the main. then we open de localhost on the navegator.
