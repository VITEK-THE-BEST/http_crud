package main

import (
	"fmt"
	"log"
	"my-crud/iternal/models"
	"my-crud/iternal/repository"
)

func main() {
	fmt.Println("Hello World!")
	users := []models.User{
		{
			Name:     "Vitek",
			Age:      10,
			Email:    "test@mail.com",
			Password: "vitek",
		}, {
			Name:     "root",
			Age:      100,
			Email:    "root@mail.com",
			Password: "root",
		},
	}
	db := repository.NewPostgresBase()
	for _, u := range users {
		err := db.CreateUser(u)
		if err != nil {
			log.Fatalln(err)
		}
	}
	users, err := db.GetUsers()
	if err != nil {
		log.Fatalln("main: ", err)
	}
	for _, u := range users {
		fmt.Printf("%T: %+v \n", u, u)
	}
}
