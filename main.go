package main

import (
	"api"
	"fmt"
)

func main() {
	users := api.GetUsers()

	fmt.Println("Users:", users)
}
