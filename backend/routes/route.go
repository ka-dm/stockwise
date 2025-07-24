package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"Stock/models"
	"Stock/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
func (h *StockHandler) FetchStocksAPI(c *gin.Context) {
	// Cargar variables de entorno (solo necesario en desarrollo)
	_ = godotenv.Load()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Obtén el token de la variable de entorno
	token := os.Getenv("API_AUTH_TOKEN")
	if token == "" {
		log.Println("API_AUTH_TOKEN not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API_AUTH_TOKEN not set"})
		return
	}

	baseURL := "https://api.karenai.click/swechallenge/list"
	nextPage := ""
	var allStocks []models.Stock

	for {
		// Construir la URL con next_page si corresponde
		url := baseURL
		if nextPage != "" {
			url = fmt.Sprintf("%s?next_page=%s", baseURL, nextPage)
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error making request: %v", err)})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("API returned status: %d", resp.StatusCode)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("API returned status: %d", resp.StatusCode)})
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
			return
		}

		// Decodificar como objeto genérico para extraer stocks y next_page
		var genericResponse map[string]interface{}
		if err := json.Unmarshal(body, &genericResponse); err != nil {
			log.Printf("Could not decode response: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode API response", "raw_response": string(body)})
			return
		}

		// Buscar el array de stocks
		var stocks []models.Stock
		for _, value := range genericResponse {
			if array, ok := value.([]interface{}); ok && len(array) > 0 {
				jsonData, _ := json.Marshal(array)
				if err := json.Unmarshal(jsonData, &stocks); err == nil {
					allStocks = append(allStocks, stocks...)
					break
				}
			}
		}

		// Buscar el next_page
		nextPage = ""
		if np, ok := genericResponse["next_page"]; ok {
			if npStr, ok := np.(string); ok && npStr != "" {
				nextPage = npStr
			}
		}

		// Si no hay next_page, salir del bucle
		if nextPage == "" {
			break
		}
	}

	// Guardar todos los stocks en la base de datos
	if len(allStocks) > 0 {
		if err := h.stockService.CreateStocksBatch(allStocks); err != nil {
			log.Printf("Error saving to database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving to database"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stocks guardados correctamente",
		"total":   len(allStocks),
	})
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API working",
		})
	})
	r.GET("/stocks", handler.GetStocksFromDB)
	r.GET("/stocks/fetch", handler.FetchStocksAPI)
}
