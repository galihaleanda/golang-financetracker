package handlers

import (
	"finance-tracker/models"
	"finance-tracker/utils"

	"github.com/gin-gonic/gin"
)

type RegisteredInput struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisteredInput
	if err := c.ShouldBindJSON(&input)
	err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Check if the email is already registered
	var existing models.User
	if result := config.DB.Where("email = ?", input.Email).First(&existing)
	result.Error == nil {
		utils.ErrorResponse(c, 400, "Email already registered")
		return
	}

	
}