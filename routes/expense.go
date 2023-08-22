package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseRoute(r *gin.RouterGroup) {
	r.GET("/expenses/:id", controller.GetExpense)
	r.GET("/expenses", controller.GetAllUserExpenses)
	r.POST("/expenses", controller.CreateExpense)
	r.PUT("/expenses/:id", controller.UpdateExpense)
	r.DELETE("/expenses/:id", controller.DeleteExpense)

}
