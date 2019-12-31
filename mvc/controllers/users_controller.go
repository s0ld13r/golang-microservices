package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/s0ld13r/golang-microservices/introduction/mvc/utils"

	"github.com/s0ld13r/golang-microservices/introduction/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_requet",
		}
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonValue)
		return
	}
	log.Printf("about to process user_id %d", userId)
	if user, apiErr := services.GetUser(userId); apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		jsonValue, _ := json.Marshal(apiErr)
		resp.Write([]byte(jsonValue))
	} else {
		jsonValue, _ := json.Marshal(user)
		resp.Write(jsonValue)
	}

}
