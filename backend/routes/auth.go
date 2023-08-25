package routes

import (
	"github.com/amikus123/cash-flowy/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/auth/login", controller.Login)
	r.POST("/auth/register", controller.Register)
	r.GET("/auth/me", controller.Me)

}
