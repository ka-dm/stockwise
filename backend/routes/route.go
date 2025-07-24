package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"Stock/models"
	"Stock/services"

	"github.com/gin-gonic/gin"
)

// StockHandler maneja las rutas relacionadas con stocks
type StockHandler struct {
	stockService *services.StockService
}

// NewStockHandler crea una nueva instancia del handler
func NewStockHandler() *StockHandler {
	return &StockHandler{
		stockService: services.NewStockService(),
	}
}

// GetStocks maneja la consulta de los stocks desde la API externa y los guarda en la BD
func (h *StockHandler) GetStocks(c *gin.Context) {
	// Crear cliente HTTP con timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Crear la petición HTTP
	req, err := http.NewRequest("GET", "https://api.karenai.click/swechallenge/list", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Agregar headers necesarios
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MjMsImVtYWlsIjoia2V2aW4uZG9tLm1vbEBnbWFpbC5jb20iLCJleHAiOjE3NTMzNzE1NzgsImlkIjoiIiwicGFzc3dvcmQiOiJ4LyoqL0ZST00vKiovdXNlcnM7LS0gLScifQ.yQ6citCrEubD6_3pt2tHgWhyv0BbiNA3jBQvVkyKmQ8")
	req.Header.Set("Content-Type", "application/json")

	// Realizar la petición
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error making request: %v", err)})
		return
	}
	defer resp.Body.Close()

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned status: %d", resp.StatusCode)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("API returned status: %d", resp.StatusCode)})
		return
	}

	// Leer el cuerpo de la respuesta para debug
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	log.Printf("API Response: %s", string(body))

	// Intentar decodificar como array directo de Stock
	var stocks []models.Stock
	if err := json.Unmarshal(body, &stocks); err == nil && len(stocks) > 0 {
		log.Printf("Successfully decoded as direct Stock array with %d items", len(stocks))

		// Usar el servicio para guardar
		if err := h.stockService.CreateStocksBatch(stocks); err != nil {
			log.Printf("Error saving to database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving to database"})
			return
		}

		c.JSON(http.StatusOK, stocks)
		return
	}

	// Si no funciona, intentar decodificar como objeto genérico
	var genericResponse map[string]interface{}
	if err := json.Unmarshal(body, &genericResponse); err == nil {
		log.Printf("Decoded as generic response: %+v", genericResponse)
		// Buscar cualquier campo que contenga un array
		for key, value := range genericResponse {
			if array, ok := value.([]interface{}); ok && len(array) > 0 {
				log.Printf("Found array in field '%s' with %d items", key, len(array))

				// Convertir el array genérico a Stock
				jsonData, _ := json.Marshal(array)
				if err := json.Unmarshal(jsonData, &stocks); err == nil {
					// Guardar en la base de datos
					if err := h.stockService.CreateStocksBatch(stocks); err != nil {
						log.Printf("Error saving to database: %v", err)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving to database"})
						return
					}

					c.JSON(http.StatusOK, stocks)
					return
				}
			}
		}
	}

	log.Printf("Could not decode response in any expected format")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode API response", "raw_response": string(body)})
}

// GetStocksFromDB obtiene los stocks desde la base de datos
func (h *StockHandler) GetStocksFromDB(c *gin.Context) {
	log.Println("Handler: Starting GetStocksFromDB")

	stocks, err := h.stockService.GetAllStocks()
	if err != nil {
		log.Printf("Handler error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching from database: %v", err)})
		return
	}

	log.Printf("Handler: Successfully retrieved %d stocks", len(stocks))
	c.JSON(http.StatusOK, stocks)
}

// RegisterStockRoutes registra las rutas relacionadas con stocks
func RegisterStockRoutes(r *gin.Engine) {
	handler := NewStockHandler()

	r.GET("/stocks", handler.GetStocks)
	r.GET("/stocks/db", handler.GetStocksFromDB)
}
