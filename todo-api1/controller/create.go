package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/model"
	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//take some data
			//save it!
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTODO(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
