package controller

import (
	"net/http"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/model"
	"github.com/gin-gonic/gin"
)

type CreateExpenseBody struct {
	Description string  `json:"description" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	CategoryID  uint    `json:"categoryId" binding:"required"`
}

func GetExpenses(c *gin.Context) {
	expenses := []model.Expense{}
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

	expense := model.Expense{
		Description: body.Description,
		Amount:      body.Amount,
		CategoryID:  body.CategoryID,
	}
	config.DB.Create(&expense)
	c.JSON(http.StatusCreated, &expense)
}
