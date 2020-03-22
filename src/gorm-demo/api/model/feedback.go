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
