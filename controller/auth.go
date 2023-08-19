package controller

import (
	"net/http"

	"github.com/amikus123/cash-flowy/model"
	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Regiter(c *gin.Context) {
	body := &RegisterBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{Username: body.Username, Email: body.Email, Password: body.Password}

	_, err := u.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
