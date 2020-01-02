package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s0ld13r/golang-microservices/mvc/utils"
)

var (
	users = map[uint64]*User{
		123: &User{
			Id: 123, FirstName: "Fede", LastName: "Leon", Email: "email@gmail.com",
		},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDao struct{}

type userDaoInterface interface {
	GetUser(uint64) (*User, *utils.ApplicationError)
}

func (u *userDao) GetUser(userId uint64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %d does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
