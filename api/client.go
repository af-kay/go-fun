package api

import "gopkg.in/h2non/gentleman.v2"

type ApiClient struct {
	*gentleman.Client
}

type ApiResponse struct {
	*gentleman.Response
}

var client *ApiClient

func init() {
	initClient()
}

func initClient() {
	if client != nil {
		return
	}

	base := "https://jsonplaceholder.typicode.com"

	client = &ApiClient{gentleman.New()}
	client.URL(base)
}

func (cli *ApiClient) fetch(path string) (*ApiResponse, error) {
	req := cli.Request()
	req.Path(path)

	response, err := req.Send()

	return &ApiResponse{response}, err
}
