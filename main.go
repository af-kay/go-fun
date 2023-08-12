package main

import (
	"api"
	"fmt"
)

func main() {
	cli := api.NewClient()
	users := api.GetUsers(cli)

	fmt.Println("Users:", users)
}
