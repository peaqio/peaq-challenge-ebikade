package models

import (
	"time"
)

// BaseModel  definition, including fields `ID`,
// `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in models
// This is for IDs that are integer
//    type User struct {
//      BaseModel
//    }
type BaseModel struct {
	ID        uint       `json:"id" gorm:"type:int(15) unsigned auto_increment;not null;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
