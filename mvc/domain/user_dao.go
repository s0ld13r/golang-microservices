package domain

import (
	"fmt"
	"net/http"

	"github.com/s0ld13r/golang-microservices/introduction/mvc/utils"
)

var users = map[uint64]*User{
	123: &User{
		Id: 123, FirstName: "Fede", LastName: "Leon", Email: "email@gmail.com",
	},
}

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprint("user", userId, "was not found"),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
