package entities

import (
	"time"

	"github.com/google/uuid"
)

type Meal struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Day       time.Time `json:"day" validate:"required" gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid4"`
	Product   []Product `json:"data_product" gorm:"json"`
	Food      []Food    `json:"data_food" gorm:"json"`
	MealType  MealType  `json:"meal_type" validate:"required,eq=breakfast|eq=lunch|eq=dinner|eq=snack"`
	Type      string    `json:"type" validate:"required,eq=product|eq=food"`
}

type MealType string

const (
	BREAKFAST MealType = "breakfast"
	LUNCH     MealType = "lunch"
	DINNER    MealType = "dinner"
	SNACK     MealType = "snack"
)

type CreateMeal struct {
	Day      time.Time `json:"day" validate:"required"`
	UserId   uuid.UUID `json:"user_id" validate:"required,uuid4"`
	Product  []Product `json:"data_product" gorm:"json"`
	Food     []Food    `json:"data_food" gorm:"json"`
	MealType MealType  `json:"meal_type" validate:"required,eq=breakfast|eq=lunch|eq=dinner|eq=snack"`
	Type     string    `json:"type" validate:"required,eq=product|eq=food"`
}
