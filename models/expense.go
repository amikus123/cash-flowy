package models

import (
	"time"
)

type Expense struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	CategoryID  uint            `json:"category_id"`
	Category    ExpenseCategory `json:"category"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
