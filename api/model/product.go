package model

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id,omitempty"`
	Name      string     `json:"name"`
	Price     int        `json:"price"`
	Qty       int        `json:"qty"`
	CreatedAt *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"type:DATETIME;default:NULL"`
}
