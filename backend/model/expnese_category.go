package model

import "github.com/amikus123/cash-flowy/config"

type ExpenseCategory struct {
	ID   uint   `json:"id"  gorm:"primaryKey"`
	Name string `json:"name"`
}

func (e *ExpenseCategory) Save() (*ExpenseCategory, error) {
	if err := config.DB.Create(&e).Error; err != nil {
		return &ExpenseCategory{}, err
	}

	return e, nil
}
