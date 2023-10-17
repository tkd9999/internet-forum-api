package main

import (
	"fmt"

	"github.com/junshintakeda/internet-forum/db"
	"github.com/junshintakeda/internet-forum/models"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})

	fmt.Println("Migrated")
}
