package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserSurvey struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	UserID    int64          `json:"user_id" validate:"required"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index,default:NULL"`
	Data      SurveyData     `json:"data" validate:"required" gorm:"json"`
}

// Target represents different target goals for a user.
type Target string

type Activity string

type SurveyData struct {
	Gender     string    `json:"gender" validate:"required,oneof=MALE FEMALE"`
	Target     Target    `json:"target" validate:"required,oneof=LOSE_WEIGHT MAINTAIN_WEIGHT GAIN_WEIGHT"`
	Growth     int       `json:"growth" validate:"required,min=0"` // Ensure growth is a positive number
	Birthday   time.Time `json:"birthday" validate:"required"`     // Use time.Time for birthday
	Activity   Activity  `json:"activity" validate:"required,oneof=SEDENTARY_LIFESTYLE MODERATE_LIFESTYLE ACTIVE_LIFESTYLE HIGHLY_ACTIVE_LIFESTYLE"`
	Weight     float32   `json:"weight" validate:"required,min=0"` // Ensure weight is a positive number
	WaistGirth float32   `json:"waist_girth" validate:"required"`  // обхват талии
	HipGirth   float32   `json:"hip_girth" validate:"required"`    // обхват бедер
}
