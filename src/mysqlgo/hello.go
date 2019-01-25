package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

func main()  {
	db, err := sql.Open("mysql","root:123456@/test")
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(db)
}