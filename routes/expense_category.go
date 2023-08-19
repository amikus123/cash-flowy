package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func ExpenseCategoryRoute(router *gin.Engine) {
	router.POST("/expnese-categories", controller.CreateExpenseCategory)

}
