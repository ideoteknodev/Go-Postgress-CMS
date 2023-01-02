package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type User struct {
	 Id     int    `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Desc   string `json:"desc"`
}


func main(){
	dbURL := "postgres://postgres:postgrespw@localhost:55000/go_cms"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

if err != nil {
	panic("failed to connect database")
}
// Auto Migrate
db.AutoMigrate(&User{})

	message := "harry"
	fmt.Println(message)
}