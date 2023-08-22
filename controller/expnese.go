package controller

import (
	"net/http"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/util/auth"
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

	var err error

	body := &CreateExpenseBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := model.Expense{
		Description: body.Description,
		Amount:      body.Amount,
		CategoryID:  body.CategoryID,
		UserID:      u.ID,
	}
	_, err = expense.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expense)

}

func GetAllUserExpenses(c *gin.Context) {

	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	expenses, err := model.GetAllExpensesByUserID(u.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, expenses)
}

func UpdateExpense(c *gin.Context) {

}

func DeleteExpense(c *gin.Context) {

}
