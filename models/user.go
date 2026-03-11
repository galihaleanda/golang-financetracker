package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	Password     string `gorm:"not null" json:"-"`
	Avatar       string `json:"avatar"`
	GoogleID     string `gorm:"uniqueIndex" json:"-"`
	IsGoogleAuth bool   `gorm:"default:false" json:"is_google_auth"`

	// User Relationship
	Categories []Category `gorm:"foreignKey:UserID"
	json:"categories,omitempty"`

	Transactions []Transaction `gorm:"foreignKey:UserID"
	json:"transactions,omitempty"`
}

// For Api Response with no password
type UserResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	IsGoogleAuth bool   `json:"is_google_auth"`
}
