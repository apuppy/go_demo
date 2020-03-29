package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostFeedback post feedback
func PostFeedback(w http.ResponseWriter, r *http.Request) {
	body := util.BodyParser(r)
	var feedback model.Feedback
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = model.NewFeedback(feedback)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, "feedback successfully added", http.StatusCreated)
}

//GetFeedbacks get rows from feedback
func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := model.GetAll(model.FEEDBACKS)
	util.ToJSON(w, feedbacks, http.StatusOK)
}

//GetFeedback get one row from feedback
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	feedback := model.GetByID(model.FEEDBACKS, id)
	util.ToJSON(w, feedback, http.StatusOK)
}

//DeleteFeedback delete feedback
func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := model.DeleteByID(model.FEEDBACKS, id)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//PutFeedback user route
func PutFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := util.BodyParser(r)
	var feedback model.Feedback
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	feedback.ID = id
	rows, err := model.UpdateFeedback(feedback)
	if err != nil {
		util.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	util.ToJSON(w, rows, http.StatusOK)
}
