package storage

import (
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
)

// RateStorage ...
type RateStorage struct {
	*EDatabase
}

// NewRateStorage Initialize Rate Storage
func NewRateStorage(db *EDatabase) *RateStorage {
	return &RateStorage{db}
}

// Fetch rate using user_id and rate_id
func (r *RateStorage) Fetch(rateID uint) *models.Rate {
	rate := models.Rate{}
	// Select resource from database
	err := r.db.Where("id=?", rateID).
		First(&rate).Error

	if rate.ID < 1 || err != nil {
		return nil
	}

	return &rate
}

// FetchByMarketName ...
func (r *RateStorage) FetchByMarketName(marketName string) *models.Rate {
	rate := models.Rate{}
	// Select resource from database
	err := r.db.
		Where("market_name=?", marketName).
		Order("created_at desc").
		Limit(1).
		First(&rate).Error

	if rate.ID < 1 || err != nil {
		return nil
	}

	return &rate
}

// Store Add a new rate
func (r *RateStorage) Store(p models.Rate) (*models.Rate, error) {

	rate := p

	err := r.db.Create(&rate).Error

	if err != nil {
		return nil, err
	}
	return r.FetchByMarketName(rate.MarketName), nil
}

// Update a rate
func (r *RateStorage) Update(rate *models.Rate) (*models.Rate, error) {

	err := r.db.Save(&rate).Error

	if err != nil {
		return nil, err
	}

	return rate, nil
}
