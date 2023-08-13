package api

import (
	"encoding/json"
	"fmt"
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

func GetUsers() (users []UserModel, err error) {
	res, err := client.fetch("/users")
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Bytes(), &users)

	return
}

func GetUser(id int) (user UserModel, err error) {
	res, err := client.fetch(fmt.Sprintf("/users/%d", id))
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Bytes(), &user)

	return
}
