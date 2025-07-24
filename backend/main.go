package main

import (
	"log"
	"time"

	"Stock/database"
	"Stock/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Reintentos para conectar a la base de datos
	var err error
	for i := 0; i < 5; i++ {
		err = database.InitDB()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/5): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database after 5 attempts:", err)
	}

	// Ejecutar migraciones
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()
	routes.RegisterStockRoutes(r)
	r.Run(":8000")
}
