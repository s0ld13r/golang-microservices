package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/s0ld13r/golang-microservices/mvc/domain"
	"github.com/s0ld13r/golang-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId uint64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func init() {
	domain.UserDao = &usersDaoMock{}
}
func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
			Code:       "not_found",
		}
	}
	const userId = 0
	user, err := UsersService.GetUser(userId)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, fmt.Sprintf("user %d does not exists", userId), err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	const userId = 123
	getUserFunction = func(userId uint64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: userId,
		}, nil
	}
	user, err := UsersService.GetUser(userId)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userId, user.Id)
}
