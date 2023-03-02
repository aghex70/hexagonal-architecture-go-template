package templates

var ResponsesTemplate = `package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONOKResponse struct {
	StatusCode   int         ` + "`json:\"statusCode\"`" + `
	Data 	   	 interface{} ` + "`json:\"data\"`" + `
}

type JSONCreatedResponse struct {
	StatusCode   int         ` + "`json:\"statusCode\"`" + `
}

type JSONErrorResponse struct {
	StatusCode   int         ` + "`json:\"statusCode\"`" + `
	Message      string      ` + "`json:\"message\"`" + `
	Data 	   	 interface{} ` + "`json:\"data\"`" + `
}

func OKResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JSONOKResponse{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func CreatedResponse(c *gin.Context) {
	c.JSON(http.StatusCreated, JSONCreatedResponse{
		StatusCode: http.StatusCreated,
	})
}

func BadRequestResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusBadRequest, JSONErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       data,
	})
}

func NotFoundResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusNotFound, JSONErrorResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Data:       data,
	})
}

func ServerErrorResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusInternalServerError, JSONErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Data:       data,
	})
}
`

func GetResponsesFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        ResponsesTemplate,
			TemplateContext: tc,
		},
	}
}
