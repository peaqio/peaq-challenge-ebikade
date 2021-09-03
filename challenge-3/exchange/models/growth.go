package models

// GrowthRecord ...
type GrowthRecord struct {
	BaseModel
	FromRateID   uint    `json:"from_rate_id" gorm:"not null;type:int(15) unsigned"`
	ToRateID     uint    `json:"to_rate_id" gorm:"not null;type:int(15) unsigned"`
	VolumeGrowth float64 `json:"volume_growth" gorm:"type:float(15,4);default:0.0"`
	HighGrowth   float64 `json:"high_growth" gorm:"type:float(15,4);default:0.0"`
	LowGrowth    float64 `json:"low_growth" gorm:"type:float(15,4);default:0.0"`
	FromDate     int64   `json:"from_date"`
	ToDate       int64   `json:"to_date"`
	FromRate     Rate    `json:"from_rate"`
	ToRate       Rate    `json:"to_rate"`
}
