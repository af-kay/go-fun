package main

import (
	"api"
	"fmt"
)

func main() {
	users, err := api.GetUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println("Users:", users)
}
