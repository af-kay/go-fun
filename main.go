package main

import (
	"api"
	"fmt"
)

func main() {
	cli := api.NewClient()

	res, err := api.Fetch(cli, "/users")
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Body: %s", res.String())
}
