package model

import (
	"time"

	"github.com/amikus123/cash-flowy/config"
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model

	ID          uint            `gorm:"primaryKey" json:"id"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	CategoryID  uint            `json:"categoryId"`
	Category    ExpenseCategory `json:"category"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

func (e *Expense) Save() (*Expense, error) {
	err := config.DB.Create(&e).Error
	if err != nil {
		return &Expense{}, err
	}
	return e, nil
}
