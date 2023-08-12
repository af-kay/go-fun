package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
)

func makeClient() *gentleman.Client {
	base := "https://jsonplaceholder.typicode.com"

	cli := gentleman.New()
	cli.URL(base)

	return cli
}

type Request struct {
	client  *gentleman.Client
	path    string
	headers map[string]string
}

func fetch(request Request) (*gentleman.Response, error) {
	req := request.client.Request()

	req.Path(request.path)
	req.SetHeaders(request.headers)

	return req.Send()
}

func main() {
	cli := makeClient()

	res, err := fetch(Request{
		client:  cli,
		path:    "/users",
		headers: map[string]string{"Client": "gentleman"},
	})
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
