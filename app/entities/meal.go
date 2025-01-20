package entities

import (
	"time"

	"gorm.io/gorm"
)

type Meal struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index,default:NULL"`
	Day       time.Time      `json:"day" validate:"required" gorm:"default:CURRENT_TIMESTAMP;index"`
	UserID    int64          `json:"user_id" validate:"required"`
	MealFoods []MealFood     `json:"meal_foods" gorm:"json"`
	MealType  MealType       `json:"meal_type" validate:"required,eq=breakfast|eq=lunch|eq=dinner|eq=snack" gorm:"index"`
	Weight    float64        `json:"weight"`
	Info      ProductStat    `json:"info" gorm:"json"`
}

type MealFood struct {
	Type   string      `json:"type" validate:"required,eq=product|eq=food"`
	Id     *int64      `json:"id"`
	Weight float64     `json:"weight" validate:"required"`
	Info   ProductStat `json:"info" gorm:"json"`
	Name   *string     `json:"name"`
}
type MealType string

const (
	BREAKFAST MealType = "breakfast"
	LUNCH     MealType = "lunch"
	DINNER    MealType = "dinner"
	SNACK     MealType = "snack"
)

type CreateMeal struct {
	Day      string     `json:"day" validate:"required"`
	Food     []MealFood `json:"food"`
	MealType MealType   `json:"meal_type" validate:"required,eq=breakfast|eq=lunch|eq=dinner|eq=snack"`
}
