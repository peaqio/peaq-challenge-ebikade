package rate

import (
	"errors"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
)

// Service provides rate operations
type Service interface {
	GetRate(uint) *models.Rate
	GetByMarketName(string) *models.Rate
	CreateRate(models.Rate) (*models.Rate, error)
	UpdateRate(*models.Rate) (*models.Rate, error)
}

type service struct {
	repo DBRepository
}

// NewService creates o rate service with the necessary dependencies
func NewService(
	repo DBRepository,
) Service {
	return &service{repo}
}

func (s *service) GetRate(rateID uint) *models.Rate {
	return s.repo.Fetch(rateID)
}

/*
* Get all rates
* @param from => the the from timestamp number to return
* @param to => to timestamp to return
 */
func (s *service) GetByMarketName(name string) *models.Rate {
	return s.repo.FetchByMarketName(name)
}

// Create New rate
func (s *service) CreateRate(r models.Rate) (*models.Rate, error) {
	// Generate ID

	rate, err := s.repo.Store(r)

	if err != nil {
		return rate, errors.New("an Error Occurred while creating reasource")
	}

	return rate, nil
}

// update existing rate
func (s *service) UpdateRate(r *models.Rate) (*models.Rate, error) {
	rate, err := s.repo.Update(r)

	if err != nil {
		return rate, errors.New("an Error Occurred while creating reasource")
	}

	return rate, nil
}
