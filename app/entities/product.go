package entities

import (
	"time"
)

type ProductParesed struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name         string    `json:"name" validate:"required" gorm:"unique"`
	Url          string    `json:"url" validate:"required" gorm:"unique"`
	CountPage    int32     `json:"count_page" validate:"required"`
	ProductCount int32     `json:"product_count" validate:"required"`
}
