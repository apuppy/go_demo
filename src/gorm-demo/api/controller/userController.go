package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostUser add user
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

//GetUsers get rows from user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetAll(model.USERS)
	util.ToJSON(w, users, http.StatusOK)
}

//GetUser get user by id request
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	user := model.GetByID(model.USERS, id)
	util.ToJSON(w, user, http.StatusOK)
}

//DeleteUser delete user by id request
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := model.DeleteByID(model.USERS, id)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusOK)
}

//PutUser update user route
func PutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := util.BodyParser(r)
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user.ID = uint32(id)
	rows, err := model.UpdateUser(user)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusOK)
}
