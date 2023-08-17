package models

type ExpenseCategory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
