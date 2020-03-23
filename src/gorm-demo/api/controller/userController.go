package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
)

// PostUser json post user
func PostUser(w http.ResponseWriter, r *http.Request) {
	body := util.BodyParser(r)
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = model.NewUser(user)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, "user successfully added", http.StatusCreated)
}
