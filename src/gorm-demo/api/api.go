package api

import "gorm-demo/api/model"

// Run api entry
func Run() {
	// db := model.Connect()
	// defer db.Close()
	// fmt.Println("Database connected!")

	// table migration
	model.AutoMigrations()
}
