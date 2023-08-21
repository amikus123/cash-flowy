package auth

import (
	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/util/token"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	u := model.User{}

	err := config.DB.Model(model.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
