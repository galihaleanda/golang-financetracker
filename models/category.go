package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Icon   string `json:"icon"`
	Color  string `json:"color"`
	UserID uint   `gorm:"not null" json:"user_id"`

	// Relationship
	Transactions []Transaction `gorm:"foreignKey:CategoryID"
	json:"transactions, omitempty"`
}

// Default Categories
var DefaultCategories = []struct {
	Name  string
	Icon  string
	Color string
}{
	{"Food & Drink", "🍔", "#FF6B6B"},
	{"Transport", "🚗", "#4ECDC4"},
	{"Shopping", "🛍️", "#45B7D1"},
	{"Bills", "💡", "#96CEB4"},
	{"Entertainment", "🎬", "#FFEAA7"},
	{"Health", "❤️", "#DDA0DD"},
	{"Education", "📚", "#98D8C8"},
	{"Salary", "💼", "#77DD77"},
}
