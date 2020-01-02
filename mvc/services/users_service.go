package services

import (
	"github.com/s0ld13r/golang-microservices/mvc/domain"
	"github.com/s0ld13r/golang-microservices/mvc/utils"
)

type usersService struct{}

var UsersService = usersService{}

func (u *usersService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
