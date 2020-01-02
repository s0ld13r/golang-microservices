package services

import (
	"net/http"

	"github.com/s0ld13r/golang-microservices/mvc/domain"
	"github.com/s0ld13r/golang-microservices/mvc/utils"
)

type itemsService struct{}

var ItemsService = itemsService{}

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
