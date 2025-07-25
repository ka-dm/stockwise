package main

import (
	"log"
	"time"

	"Stock/database"
	"Stock/routes"

	"github.com/gin-contrib/cors"
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

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.12.149:5173", "http://localhost:5173"}, // URL de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterStockRoutes(r)
	r.Run(":8000")
}
