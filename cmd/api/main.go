package main

import (
	"github.com/basic-go/internal/book"
	"github.com/basic-go/internal/database"
	"github.com/basic-go/internal/router"
)

func main() {

	db := database.NewSQLiteConnection()

	db.AutoMigrate(&book.Book{})

	repo := book.NewRepository(db)

	service := book.NewService(repo)

	handler := book.NewHandler(service)

	r := router.SetupRouter(handler)

	r.Run(":8080")
}
