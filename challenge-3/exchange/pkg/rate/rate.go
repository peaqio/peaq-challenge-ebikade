package rate

import (
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
)

// DBRepository ...
type DBRepository interface {
	Fetch(uint) *models.Rate
	FetchByMarketName(string) *models.Rate
	Store(models.Rate) (*models.Rate, error)
	Update(*models.Rate) (*models.Rate, error)
}
