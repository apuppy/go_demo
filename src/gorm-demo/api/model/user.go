package model

import (
	"gorm-demo/api/security"
	"time"
)

// User user table
type User struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string     `gorm:"type:varchar(20);not null;unique_index" json:"nickname"`
	Email     string     `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	Password  string     `gorm:"type:varchar(60);not null" json:"password"`
	CreatedAt time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post     `gorm:"ForeignKey:UserID" json:"posts"`
	Feedbacks []Feedback `gorm:"ForeignKey:UserID" json:"feedbacks"`
}

// NewUser add new user
func NewUser(user User) error {
	db := Connect()
	defer db.Close()
	var err error
	user.Password, err = security.Hash(user.Password)
	if err != nil {
		return nil
	}
	err = db.Create(&user).Error
	return err
}
