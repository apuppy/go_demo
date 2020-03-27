package model

import "time"

// Post post
type Post struct {
	ID          uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Description string     `gorm:"type:varchar(255)" json:"description"`
	ImageURL    string     `gorm:"type:varchar(255)" json:"image_url"`
	Subtitle    string     `gorm:"type:varchar(100)" json:"subtitle"`
	UserID      uint32     `gorm:"not null" json:"user_id"`
	User        User       `json:"user"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	Feedbacks   []Feedback `gorm:"ForeignKey:PostID" json:"feedbacks"`
}

// NewPost add new post
func NewPost(post Post) error {
	db := Connect()
	defer db.Close()
	return db.Create(&post).Error
}

//GetPosts get post rows
func GetPosts() []Post {
	db := Connect()
	defer db.Close()
	var posts []Post
	db.Order("id ASC").Find(&posts)
	for i := range posts {
		db.Model(&posts[i]).Related(&posts[i].User)
		posts[i].Feedbacks = GetFeedbackByPost(posts[i])
	}
	return posts
}
