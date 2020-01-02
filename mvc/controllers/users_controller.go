package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/s0ld13r/golang-microservices/mvc/utils"

	"github.com/s0ld13r/golang-microservices/mvc/services"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_requet",
		}
		utils.RespondError(c, apiError)
		return
	}
	log.Printf("about to process user_id %d", userId)
	if user, apiErr := services.UsersService.GetUser(userId); apiErr != nil {
		utils.RespondError(c, apiErr)
	} else {
		utils.Resspond(c, http.StatusOK, user)
	}

}
