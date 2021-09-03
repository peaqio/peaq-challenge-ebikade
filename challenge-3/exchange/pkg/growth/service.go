package growth

import (
	"errors"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
)

// Service provides rate operations
type Service interface {
	GetGrowthRecord(uint) *models.GrowthRecord
	GetByTimestamp(int64, int64) ([]*models.GrowthRecord, error)
	CreateGrowthRecord(models.GrowthRecord) (*models.GrowthRecord, error)
	UpdateGrowthRecord(*models.GrowthRecord) (*models.GrowthRecord, error)
}

type service struct {
	repo DBRepository
}

// NewService creates rate service with the necessary dependencies
func NewService(
	repo DBRepository,
) Service {
	return &service{repo}
}

func (s *service) GetGrowthRecord(rateID uint) *models.GrowthRecord {
	return s.repo.Fetch(rateID)
}

/*
* Get all rates
* @param from => the the from timestamp number to return
* @param to => to timestamp to return
 */
func (s *service) GetByTimestamp(from, to int64) ([]*models.GrowthRecord, error) {
	return s.repo.FetchByTimestamp(from, to), nil
}

// Create New rate
func (s *service) CreateGrowthRecord(r models.GrowthRecord) (*models.GrowthRecord, error) {
	// Generate ID

	rate, err := s.repo.Store(r)

	if err != nil {
		return rate, errors.New("an Error Occurred while creating reasource")
	}

	return rate, nil
}

// update existing rate
func (s *service) UpdateGrowthRecord(r *models.GrowthRecord) (*models.GrowthRecord, error) {
	rate, err := s.repo.Update(r)

	if err != nil {
		return rate, errors.New("an Error Occurred while creating reasource")
	}

	return rate, nil
}
