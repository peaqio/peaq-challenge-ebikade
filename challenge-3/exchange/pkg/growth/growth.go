package growth

import (
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
)

// DBRepository ...
type DBRepository interface {
	Fetch(uint) *models.GrowthRecord
	FetchByTimestamp(int64, int64) []*models.GrowthRecord
	Store(models.GrowthRecord) (*models.GrowthRecord, error)
	Update(*models.GrowthRecord) (*models.GrowthRecord, error)
}
