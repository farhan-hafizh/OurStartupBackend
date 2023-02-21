package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// add meta with type meta inside response and data with type interface because
// it's dynamic object
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//create response
func CreateResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	// if data exist return with data
	if data != nil {
		return Response{
			Meta: meta,
			Data: data,
		}
	}
	// if data not exist or nil
	return Response{
		Meta: meta,
	}
}

func SendResponse(c *gin.Context, message string, code int, status string, data interface{}) {
	c.JSON(
		code,
		CreateResponse(message,
			code,
			status,
			data))
}

func SendErrorResponse(c *gin.Context, message string, code int, status string, isValidationError bool, err error) {
	var errorText []string

	errorText[0] = err.Error()

	if isValidationError {
		errorText = FormatValidationError(err)
	}

	errorResponse := gin.H{"errors": errorText}

	SendResponse(c, message, code, status, errorResponse)

	c.Error(err)
	c.Abort()
}

func FormatValidationError(err error) []string {
	var errors []string

	// loop through error that had changed to ValidationErrors type
	// then append the error to errors
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
