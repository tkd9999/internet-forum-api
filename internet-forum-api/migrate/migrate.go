package main

import (
	"fmt"

	"github.com/junshintakeda/internet-forum/db"
	"github.com/junshintakeda/internet-forum/models"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	dbConn.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})

	fmt.Println("Migrated")
}
