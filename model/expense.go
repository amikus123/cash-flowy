package model

import (
	"time"

	"github.com/amikus123/cash-flowy/config"
)

type Expense struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	CategoryID  uint            `json:"categoryId"`
	Category    ExpenseCategory `json:"category"`
	UserID      uint            `json:"userId"`
	User        User            `json:"user"`
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

func GetAllExpensesByUserID(userID uint) ([]Expense, error) {
	expenses := []Expense{}
	if err := config.DB.Find(&expenses, "user_id =?", userID).Error; err != nil {
		return []Expense{}, err
	}
	return expenses, nil

}
