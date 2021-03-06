package model

import "time"

// Feedback feedback
type Feedback struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Comment   string    `gorm:"size:255;not null" json:"comment"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	User      User      `json:"user"`
	PostID    uint32    `gorm:"not null" json:"post_id"`
	Post      Post      `json:"post"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// NewFeedback add new feedback
func NewFeedback(feedback Feedback) error {
	db := Connect()
	defer db.Close()
	return db.Create(&feedback).Error
}

//GetFeedbacks get feedback rows
func GetFeedbacks() []Feedback {
	db := Connect()
	defer db.Close()
	var feedbacks []Feedback
	db.Order("id ASC").Find(&feedbacks)
	for i := range feedbacks {
		db.Model(&feedbacks[i]).Related(&feedbacks[i].User)
	}
	return feedbacks
}

//GetFeedbackByPost get feedback by post_id
func GetFeedbackByPost(post Post) []Feedback {
	db := Connect()
	defer db.Close()
	var feedbacks []Feedback
	db.Where("post_id = ?", post.ID).Find(&feedbacks)
	for i := range feedbacks {
		db.Model(&feedbacks[i]).Related(&post)
	}
	return feedbacks
}

//UpdateFeedback update feedback
func UpdateFeedback(feedback Feedback) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&feedback).Where("id = ?", feedback.ID).UpdateColumns(
		map[string]interface{}{
			"comment": feedback.Comment,
		},
	)
	return rs.RowsAffected, rs.Error
}
