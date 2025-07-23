package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Stock representa la estructura de un stock
type Stock struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

// GetStocks maneja la consulta de los stocks
func GetStocks(c *gin.Context) {
	// Datos de ejemplo, en una aplicación real se consultaría la base de datos
	stocks := []Stock{
		{Ticker: "AAPL", TargetFrom: "100", TargetTo: "150", Company: "Apple Inc.", Action: "Buy", Brokerage: "100", RatingFrom: "1", RatingTo: "5", Time: "10:00"},
		{Ticker: "GOOGL", TargetFrom: "200", TargetTo: "250", Company: "Alphabet Inc.", Action: "Sell", Brokerage: "100", RatingFrom: "1", RatingTo: "5", Time: "10:00"},
		{Ticker: "AMZN", TargetFrom: "300", TargetTo: "350", Company: "Amazon.com Inc.", Action: "Hold", Brokerage: "100", RatingFrom: "1", RatingTo: "5", Time: "10:00"},
	}

	c.JSON(http.StatusOK, stocks)
}

// RegisterStockRoutes registra las rutas relacionadas con stocks
func RegisterStockRoutes(r *gin.Engine) {
	r.GET("/stocks", GetStocks)
}
