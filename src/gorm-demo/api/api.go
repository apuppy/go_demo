package api

import "gorm-demo/api/model"

// Run api entry
func Run() {
	// db := model.Connect()
	// defer db.Close()
	// fmt.Println("Database connected!")

	// table migration
	model.AutoMigrations()

	// add user
	model.NewUser(model.User{Nickname: "简爱", Email: "jane@gmail.com", Password: "123987"})
}
