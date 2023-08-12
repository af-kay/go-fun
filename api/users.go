package api

import (
	"encoding/json"
	"gopkg.in/h2non/gentleman.v2"
)

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		}
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func GetUsers(cli *gentleman.Client) (users []UserModel) {
	res, err := Fetch(cli, "/users")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(res.Bytes(), &users); err != nil {
		panic(err)
	}

	return
}
