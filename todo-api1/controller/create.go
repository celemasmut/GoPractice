package controller

import (
	"net/http"

	"github.com/celemasmut/GoPractice/GoPractice/todo-api1/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//take some data
			//save it!
			if err := model.CreateTODO(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}
