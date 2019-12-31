package services

import (
	"github.com/s0ld13r/golang-microservices/introduction/mvc/domain"
	"github.com/s0ld13r/golang-microservices/introduction/mvc/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
