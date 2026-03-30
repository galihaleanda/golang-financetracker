package handlers

import (
	"finance-tracker/config"
	"finance-tracker/models"
	"finance-tracker/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	// Check if the email is already registered
	var existing models.User
	if result := config.DB.Where("email = ?", input.Email).First(&existing); result.Error == nil {
		utils.ErrorResponse(c, 400, "Email already registered")
		return
	}

	// Hash password dengan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to process password")
		return
	}
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to create user")
		return
	}

	// Create default categories for new user
	createDefaultCategories(user.ID)

	token, _ := utils.GenerateToken(user.ID, user.Email)
	utils.SuccessResponse(c, 201, "Registration success", gin.H{
		"token": token,
		"user":  toUserResponse(user),
	})

}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.ErrorResponse(c, 401, "Invalid email or password")
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(input.Password)); err != nil {
			utils.ErrorResponse(c, 401, "Invalid email or password")
			return
	}

	token, _ := utils.GenerateToken(user.ID, user.Email)
	utils.SuccessResponse(c, 200, "Login Successful", gin.H{
		"token": token,
		"user": toUserResponse(user),
	})
}

// Helper functions
func toUserResponse(user models.User) models.UserResponse {
	return models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Avatar: user.Avatar,
		IsGoogleAuth: user.IsGoogleAuth,
	}
}

// Helper functions
func createDefaultCategories(userID uint) {
	for _, cat := range models.DefaultCategories {
		config.DB.Create (&models.Category{
			Name: cat.Name,
			Icon: cat.Icon,
			Color: cat.Color,
			UserID: userID,
		})
	}
}
