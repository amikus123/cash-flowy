package model

import (
	"time"
)

type Expense struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	CategoryID  uint            `json:"categoryId"`
	Category    ExpenseCategory `json:"category"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}
