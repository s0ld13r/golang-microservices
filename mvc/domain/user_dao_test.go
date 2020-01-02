package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	const id = 0
	user, err := UserDao.GetUser(id)
	assert.Nil(t, user, "We are not expecting a user with id 0 ")
	assert.NotNil(t, err, "we were expecting an error when user id = 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we are expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	const id = 123
	user, err := UserDao.GetUser(id)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, id, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "Leon", user.LastName)
	assert.EqualValues(t, "email@gmail.com", user.Email)
}
