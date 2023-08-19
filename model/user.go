package model

import (
	"html"
	"strings"
	"time"

	"github.com/amikus123/cash-flowy/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	CreatedAt time.Time
}

func (u *User) Save() (*User, error) {
	err := config.DB.Create(u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(hash)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
