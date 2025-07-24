package services

import (
	"Stock/models"
	"Stock/repositories"
	"log"
)

// StockService maneja la lógica de negocio para stocks
type StockService struct {
	repo *repositories.StockRepository
}

// NewStockService crea una nueva instancia del servicio
func NewStockService() *StockService {
	return &StockService{
		repo: repositories.NewStockRepository(),
	}
}

// CreateStock crea un nuevo stock con validaciones
func (s *StockService) CreateStock(stock *models.Stock) error {
	if err := stock.Validate(); err != nil {
		return err
	}
	return s.repo.Create(stock)
}

// CreateStocksBatch crea múltiples stocks
func (s *StockService) CreateStocksBatch(stocks []models.Stock) error {
	log.Printf("Service: Creating batch of %d stocks", len(stocks))
	for _, stock := range stocks {
		if err := stock.Validate(); err != nil {
			log.Printf("Validation error for stock %s: %v", stock.Ticker, err)
			return err
		}
	}
	return s.repo.CreateBatch(stocks)
}

// GetAllStocks obtiene todos los stocks
func (s *StockService) GetAllStocks() ([]models.Stock, error) {
	log.Println("Service: Getting all stocks")
	stocks, err := s.repo.FindAll()
	if err != nil {
		log.Printf("Service error getting stocks: %v", err)
		return nil, err
	}
	log.Printf("Service: Retrieved %d stocks", len(stocks))
	return stocks, err
}

// GetStockByTicker obtiene un stock por ticker
func (s *StockService) GetStockByTicker(ticker string) (*models.Stock, error) {
	return s.repo.FindByTicker(ticker)
}
