package controller

import (
	"net/http"
	"strconv"

	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/util/auth"
	"github.com/gin-gonic/gin"
)

type CreateExpenseBody struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	CategoryID  uint    `json:"categoryId" binding:"required"`
}

func GetExpense(c *gin.Context) {
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	expense, err := model.GetExpensesByID(uint(u64))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if expense.UserID != u.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cann't acces"})
		return

	}

	c.JSON(http.StatusOK, expense)
}

func CreateExpense(c *gin.Context) {
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
		Title:       body.Title,
		Description: body.Description,
		Amount:      body.Amount,
		CategoryID:  body.CategoryID,
		UserID:      u.ID,
	}
	_, err = expense.Create()

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
		return
	}

	expenses, err := model.GetAllExpensesByUserID(u.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func UpdateExpense(c *gin.Context) {
	body := &CreateExpenseBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	expense, err := model.GetExpensesByID(uint(u64))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if expense.UserID != u.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cann't acces"})
		return
	}
	expense.Description = body.Description
	expense.Amount = body.Amount
	expense.Title = body.Title
	expense.CategoryID = body.CategoryID
	exp, err := expense.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exp)

}

func DeleteExpense(c *gin.Context) {

	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	expense, err := model.GetExpensesByID(uint(u64))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	u, err := auth.GetUserFromContect(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if expense.UserID != u.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cann't acces"})
		return
	}

	exp, err := expense.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exp)

}
