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
	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controller.PutUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	/* posts route */
	r.HandleFunc("/posts", controller.PostPost).Methods("POST")
	r.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.PutPost).Methods("PUT")
	r.HandleFunc("/posts/{id}", controller.DeletePost).Methods("DELETE")

	/* feedback route */
	r.HandleFunc("/feedbacks", controller.PostFeedback).Methods("POST")
	r.HandleFunc("/feedbacks", controller.GetFeedbacks).Methods("GET")
	r.HandleFunc("/feedbacks/{id}", controller.GetFeedback).Methods("GET")
	r.HandleFunc("/feedbacks/{id}", controller.PutFeedback).Methods("PUT")
	r.HandleFunc("/feedbacks/{id}", controller.DeleteFeedback).Methods("DELETE")

	return r
}
