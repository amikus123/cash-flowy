package main

import (
	"github.com/gin-gonic/gin"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/model"
	"github.com/amikus123/cash-flowy/routes"
)

func main() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&model.User{}, &model.ExpenseCategory{}, &model.User{})

	router := gin.New()
	routes.ExpenseRoute(router)
	routes.ExpenseCategoryRoute(router)
	routes.AuthRoute(router)

	router.Run(":8080")
}
