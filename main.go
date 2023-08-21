package main

import (
	"github.com/gin-gonic/gin"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/middleware"
	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/routes"
)

func main() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&model.User{}, &model.ExpenseCategory{}, &model.User{})

	r := gin.New()
	public := r.Group("/api")
	routes.AuthRoute(public)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	routes.ExpenseRoute(protected)
	routes.ExpenseCategoryRoute(protected)

	r.Run(":8080")
}
