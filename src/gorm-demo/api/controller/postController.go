package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
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
