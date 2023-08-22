package controller

import (
	"net/http"

	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/util/auth"
	"github.com/amikus123/cash-flowy/util/token"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	body := &LoginBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{Email: body.Email, Password: body.Password}

	token, err := auth.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type RegisterBody struct {
	LoginBody
	Username string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	body := &RegisterBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{Username: body.Username, Email: body.Email, Password: body.Password}

	savedUser, err := u.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := token.GenerateToken(savedUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func Me(c *gin.Context) {
	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u})
}
