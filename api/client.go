package api

import "gopkg.in/h2non/gentleman.v2"

func NewClient() *gentleman.Client {
	base := "https://jsonplaceholder.typicode.com"

	cli := gentleman.New()
	cli.URL(base)

	return cli
}

func Fetch(cli *gentleman.Client, path string) (*gentleman.Response, error) {
	req := cli.Request()

	req.Path(path)

	return req.Send()
}
