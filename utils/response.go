package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success dengan respons data
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error Response
func ErrorResponse(c *gin.Context, err error) {
	c.JSON(400, APIResponse{
		Success: false,
		Message: "Validation Failed",
		Error:   err.Error(),
	})
}
