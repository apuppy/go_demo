package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostPost add post row
func PostPost(w http.ResponseWriter, r *http.Request) {
	body := util.BodyParser(r)
	var post model.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = model.NewPost(post)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, "post successfully added", http.StatusCreated)
}

//GetPosts get rows from post
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := model.GetPosts()
	util.ToJSON(w, posts, http.StatusOK)
}

//GetPost get one row
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	post := model.GetByID(model.POSTS, id)
	util.ToJSON(w, post, http.StatusOK)
}

//DeletePost delete post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := model.DeleteByID(model.POSTS, id)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//PutPost update post route
func PutPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := util.BodyParser(r)
	var post model.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	post.ID = uint32(id)
	rows, err := model.UpdatePost(post)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusCreated)
}
