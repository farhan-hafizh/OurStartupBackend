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
		errorValidation := gin.H{"errors": helper.FormatValidationError(err)}

		c.JSON(
			http.StatusBadRequest,
			helper.CreateResponse("Register account failed",
				http.StatusUnprocessableEntity,
				"Failed",
				errorValidation))
		c.Error(err)
		c.Abort()
		return
	}

	newUser, err := handler.userService.RegisterUser(*input)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helper.CreateResponse("Register account failed",
				http.StatusBadRequest, "Failed",
				err.Error()))

		c.Error(err)
		c.Abort()
		return
	}

	response :=
		helper.CreateResponse(
			"Your account successfully registered!",
			200,
			"Success",
			user.FormatRegisterResponse(newUser))

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) Login(c *gin.Context) {
	// init temp
	input := &user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorValidation := gin.H{"errors": helper.FormatValidationError(err)}

		c.JSON(
			http.StatusBadRequest,
			helper.CreateResponse("Login failed",
				http.StatusUnprocessableEntity,
				"Failed",
				errorValidation))
		c.Error(err)
		c.Abort()
		return
	}

	foundedUser, err := handler.userService.Login(*input)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			helper.CreateResponse("Login failed",
				http.StatusNotFound,
				"Failed",
				gin.H{"errors": err.Error()}))
		c.Error(err)
		c.Abort()
		return
	}
	response :=
		helper.CreateResponse("Your account successfully registered!", 200, "Success", user.FormatRegisterResponse(foundedUser))

	c.JSON(http.StatusOK, response)
}
