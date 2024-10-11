package main

import (
	"context"
	"log"
	"my-crud/iternal/models"
	"my-crud/iternal/repository"
)

func main() {
	db := repository.NewPostgresBase()
	err := db.ResetModel(context.Background(), &models.User{})
	if err != nil {
		log.Fatalln("migrations: ", err)
	}
}
