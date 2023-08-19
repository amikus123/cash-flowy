package controller

import (
	"net/http"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/model"

	"github.com/gin-gonic/gin"
)

type CreateExpenseCategoryBody struct {
	Name string `json:"name" binding:"required"`
}

func CreateExpenseCategory(c *gin.Context) {
	body := &CreateExpenseCategoryBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expenseCategory := model.ExpenseCategory{Name: body.Name}
	config.DB.Create(&expenseCategory)
	c.JSON(http.StatusCreated, &expenseCategory)
}
