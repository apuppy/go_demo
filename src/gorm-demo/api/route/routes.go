package route

import (
	"gorm-demo/api/controller"

	"github.com/gorilla/mux"
)

// NewRouter create http route
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", controller.PostUser).Methods("POST")
	r.HandleFunc("/posts", controller.PostPost).Methods("POST")
	r.HandleFunc("/feedbacks", controller.PostFeedback).Methods("POST")
	return r
}
