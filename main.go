package main

import (
	"api"
	"fmt"
)

func main() {
	fmt.Println("======================")
	users, err := api.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println("Users:", users)

	fmt.Println("======================")
	user, err := api.GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("User:", user)
}
