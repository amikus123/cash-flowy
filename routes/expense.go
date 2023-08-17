package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseRoute(router *gin.Engine) {
	router.GET("/expenses", controller.GetExpenses)
	router.GET("/expenses:id", controller.GetExpense)
	router.POST("/expenses", controller.CreateExpense)

}
