package controller

import (
	"net/http"
)

//Register es a function that create and return e new multiplaxer
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping()) //on /ping we execute the ping func.
	mux.HandleFunc("/", create())
	return mux
}
