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
	_, err := expenseCategory.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenseCategory)

}

func GetExpenseCategories(c *gin.Context) {

	expenseCatetegories := []model.ExpenseCategory{}
	res := config.DB.Find(&expenseCatetegories)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, expenseCatetegories)

}
