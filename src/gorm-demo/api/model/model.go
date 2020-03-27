package model

import "github.com/jinzhu/gorm"

// table names
const (
	USERS     = "users"
	POSTS     = "posts"
	FEEDBACKS = "feedbacks"
)

//GetAll get all data from specific table
func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()

	switch table {
	case USERS:
		return db.Order("id ASC").Find(&[]User{}).Value
	case POSTS:
		return db.Order("id ASC").Find(&[]Post{}).Value
	case FEEDBACKS:
		return db.Order("id ASC").Find(&[]Feedback{}).Value
	}
	return nil
}

// GetByID get row by id
func GetByID(table string, id uint64) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Where("id = ?", id).Find(&User{}).Value
	case POSTS:
		return db.Where("id = ?", id).Find(&Post{}).Value
	case FEEDBACKS:
		return db.Where("id = ?", id).Find(&Feedback{}).Value
	}
	return nil
}

// DeleteByID delete row by id
func DeleteByID(table string, id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case USERS:
		rs = db.Where("id = ?", id).Delete(&User{})
	case POSTS:
		rs = db.Where("id = ?", id).Delete(&Post{})
	case FEEDBACKS:
		rs = db.Where("id = ?", id).Delete(&Feedback{})
	}
	return rs.RowsAffected, rs.Error
}
