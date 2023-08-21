package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseRoute(r *gin.RouterGroup) {
	r.GET("/expenses", controller.GetExpenses)
	r.GET("/expenses:id", controller.GetExpense)
	r.POST("/expenses", controller.CreateExpense)
}
