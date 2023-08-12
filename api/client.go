package api

import "gopkg.in/h2non/gentleman.v2"

var cli *gentleman.Client

func getClient() *gentleman.Client {
	if cli != nil {
		return cli
	}

	base := "https://jsonplaceholder.typicode.com"

	cli = gentleman.New()
	cli.URL(base)

	return cli
}

func fetch(cli *gentleman.Client, path string) (*gentleman.Response, error) {
	req := cli.Request()

	req.Path(path)

	return req.Send()
}
