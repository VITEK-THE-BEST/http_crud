package main

import (
	"context"
	"log"
	"my-crud/iternal/modles"
	"my-crud/iternal/repository"
)

func main() {
	db := repository.NewPostgresBase()
	err := db.ResetModel(context.Background(), &modles.User{})
	if err != nil {
		log.Fatalln("migrations: ", err)
	}
}
