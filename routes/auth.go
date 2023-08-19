package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/login", controller.GetExpenses)
	router.POST("/auth/register", controller.Regiter)
}
