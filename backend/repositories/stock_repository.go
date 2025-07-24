package repositories

import (
	"Stock/database"
	"Stock/models"
	"log"

	"gorm.io/gorm/clause"
)

// StockRepository maneja las operaciones de base de datos para stocks
type StockRepository struct{}

// NewStockRepository crea una nueva instancia del repositorio
func NewStockRepository() *StockRepository {
	return &StockRepository{}
}

// Create crea un nuevo stock
func (r *StockRepository) Create(stock *models.Stock) error {
	log.Printf("Creating stock: %s", stock.Ticker)
	return database.DB.Create(stock).Error
}

// CreateBatch crea múltiples stocks
func (r *StockRepository) CreateBatch(stocks []models.Stock) error {
	log.Printf("Creating batch of %d stocks", len(stocks))
	return database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&stocks).Error
}

// FindAll obtiene todos los stocks
func (r *StockRepository) FindAll() ([]models.Stock, error) {
	log.Println("Fetching all stocks from database")
	var stocks []models.Stock
	err := database.DB.Find(&stocks).Error
	if err != nil {
		log.Printf("Error in FindAll: %v", err)
		return nil, err
	}
	log.Printf("Found %d stocks in database", len(stocks))
	return stocks, err
}

// FindByTicker obtiene un stock por ticker
func (r *StockRepository) FindByTicker(ticker string) (*models.Stock, error) {
	log.Printf("Fetching stock with ticker: %s", ticker)
	var stock models.Stock
	err := database.DB.Where("ticker = ?", ticker).First(&stock).Error
	if err != nil {
		log.Printf("Error in FindByTicker: %v", err)
		return nil, err
	}
	return &stock, nil
}

// Update actualiza un stock
func (r *StockRepository) Update(stock *models.Stock) error {
	return database.DB.Save(stock).Error
}

// Delete elimina un stock
func (r *StockRepository) Delete(stock *models.Stock) error {
	return database.DB.Delete(stock).Error
}
