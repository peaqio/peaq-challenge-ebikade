package models

import (
	"time"
)

type Rate struct {
	BaseModel
	MarketName string    `json:"market_name" gorm:"not null;type:varchar(20)"`
	High       float64   `json:"high" gorm:"type:float(11,8);default:0.0"`
	Low        float64   `json:"low" gorm:"type:float(11,8);default:0.0"`
	Volume     float64   `json:"volume" gorm:"type:float(16,8);default:0.0"`
	Timestamp  time.Time `json:"timestamp"`
}
