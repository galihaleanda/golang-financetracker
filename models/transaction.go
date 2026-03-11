package models

import "gorm.io/gorm"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	gorm.Model
	Title      string          `gorm:"not null" json:"title"`
	Amount     float64         `gorm:"not null;check:amount > 0" json:"amount"`
	Type       TransactionType `gorm:"not null" json:"type"`
	Note       string          `json:"note"`
	UserID     uint            `gorm:"not null;index" json:"user_id"`
	CategoryID uint            `gorm:"not null" json:"category_id"`

	// Preload Relationship
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	User     User     `gorm:"foreignKey:UserID" json:"-"`
}

// For Query Filter
type TransactionFilter struct {
	StartDate  string `form:"start_date"`
	EndDate    string `form:"end_date"`
	CategoryID uint   `form:"category_id"`
	Type       string `form:"type"`
	Search     string `form:"search"`
	Page       int    `form:"page"`
	Limit      int    `form:"page"`
}
