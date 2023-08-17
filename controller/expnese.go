package controller

import (
	"fmt"
	"net/http"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/models"
	"github.com/gin-gonic/gin"
)

type CreateExpenseBody struct {
	Description string `json:"description" binding:"required"`
}

func GetExpenses(c *gin.Context) {
	expenses := []models.Expense{}
	config.DB.Find(&expenses)
	c.JSON(200, &expenses)
}

func GetExpense(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Hello %s", id)
}

func CreateExpense(c *gin.Context) {
	body := &CreateExpenseBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(body)
	c.JSON(http.StatusCreated, &body)
}
