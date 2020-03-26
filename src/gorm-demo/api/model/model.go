package model

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
		return db.Find(&[]User{}).Value
	case POSTS:
		return db.Find(&[]Post{}).Value
	case FEEDBACKS:
		return db.Find(&[]Feedback{}).Value
	}
	return nil
}
