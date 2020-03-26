package controller

import (
	"encoding/json"
	"gorm-demo/api/model"
	"gorm-demo/api/util"
	"net/http"
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
