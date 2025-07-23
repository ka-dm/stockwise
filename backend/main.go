package main

import (
	"Stock/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterStockRoutes(r)
	r.Run(":8000")
}
