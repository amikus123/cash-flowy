package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseCategoryRoute(r *gin.RouterGroup) {
	r.POST("/expnese-categories", controller.CreateExpenseCategory)

}
