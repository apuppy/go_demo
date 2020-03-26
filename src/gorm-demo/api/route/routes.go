package route

import (
	"gorm-demo/api/controller"

	"github.com/gorilla/mux"
)

// NewRouter create http route
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	/* users route */
	r.HandleFunc("/users", controller.PostUser).Methods("POST")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")

	/* posts route */
	r.HandleFunc("/posts", controller.PostPost).Methods("POST")
	r.HandleFunc("/posts", controller.GetPosts).Methods("GET")

	/* feedback route */
	r.HandleFunc("/feedbacks", controller.PostFeedback).Methods("POST")
	r.HandleFunc("/feedbacks", controller.GetFeedbacks).Methods("GET")

	return r
}
