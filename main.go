package main

import (
	"github.com/gin-gonic/gin"

	"github.com/amikus123/cash-flowy/config"
	"github.com/amikus123/cash-flowy/routes"
)

func main() {
	config.ConnectToDB()

	router := gin.New()
	routes.ExpenseRoute(router)
	routes.ExpenseCategoryRoute(router)

	router.Run(":8080")
}
