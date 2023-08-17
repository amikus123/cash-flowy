package repositories

import "gorm.io/gorm"

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepositroy(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{db}
}

// func (* ExpenseRepository) Save (expense *models.Expense) error {
