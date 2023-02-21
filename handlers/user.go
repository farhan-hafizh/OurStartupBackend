package handlers

import (
	"net/http"
	"ourstartup/helper"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func CreateUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	// get input register struct
	input := &user.RegisterUserInput{}

	//bind the request body to input
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Register account failed",
			http.StatusUnprocessableEntity,
			"Failed",
			true,
			err)
		return
	}

	newUser, err := handler.userService.RegisterUser(*input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Register account failed",
			http.StatusBadRequest,
			"Failed",
			false,
			err)
		return
	}

	helper.SendResponse(
		c,
		"Your account successfully registered!",
		http.StatusOK,
		"Success",
		user.FormatRegisterResponse(newUser))

}

func (handler *userHandler) Login(c *gin.Context) {
	// init temp
	input := &user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)
	//if validation error
	if err != nil {
		helper.SendErrorResponse(
			c,
			"Login failed",
			http.StatusUnprocessableEntity,
			"Failed",
			true,
			err)
		return
	}

	loggedinUser, err := handler.userService.Login(*input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Login failed",
			http.StatusNotFound,
			"Failed",
			false,
			err)
		return
	}

	token := "tokentokentoken"

	helper.SendResponse(
		c,
		"You're successfully loggedin!",
		http.StatusOK,
		"Success",
		user.FormatLoginResponse(loggedinUser, token))
}

func (handler *userHandler) CheckEmailAvailability(c *gin.Context) {
	input := &user.CheckEmailInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to check",
			http.StatusUnprocessableEntity,
			"Failed",
			true,
			err)
		return
	}

	isEmailAvailable, err := handler.userService.IsEmailAvailable(*input)

	responseData := gin.H{
		"is_available": isEmailAvailable,
		"email":        input.Email,
	}

	metaMessage := "Email is registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	helper.SendResponse(c, metaMessage, http.StatusOK, "Success", responseData)
}
