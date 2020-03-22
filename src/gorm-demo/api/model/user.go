package model

import (
	"time"
)

// User user table
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"type:varchar(20);not null;unique_index" json:"nickname"`
	Email     string    `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(60);not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
