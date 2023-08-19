package model

type ExpenseCategory struct {
	ID   uint   `json:"id"  gorm:"primaryKey"`
	Name string `json:"name"`
}
