package utils

import "github.com/gin-gonic/gin"

// API Response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success Reponse w data
func SuccessResponse (c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse {
		Success: false,
		Message: message,
		Data: data,
	})
}

// Error Response 
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, APIResponse {
		Success: false,
		Message: message,
		Error: message,
	})
}

// Validation error response
func ValidationErrorResponse (c *gin.Context, err error) {
	c.JSON(400, APIResponse{
		Success: false,
		Message: "Validation failed",
		Error: err.Error(),
	})
}
