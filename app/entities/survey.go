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

const (
	LOSE_WEIGHT     Target = "LOSE_WEIGHT"     // Losing weight
	MAINTAIN_WEIGHT Target = "MAINTAIN_WEIGHT" // Maintaining weight
	GAIN_WEIGHT     Target = "GAIN_WEIGHT"     // Gaining weight
)

// Activity represents different activity levels of a user.
type Activity string

const (
	SEDENTARY_LIFESTYLE     Activity = "SEDENTARY_LIFESTYLE"     // Sedentary lifestyle
	MODERATE_LIFESTYLE      Activity = "MODERATE_LIFESTYLE"      // Moderate lifestyle
	ACTIVE_LIFESTYLE        Activity = "ACTIVE_LIFESTYLE"        // Active lifestyle
	HIGHLY_ACTIVE_LIFESTYLE Activity = "HIGHLY_ACTIVE_LIFESTYLE" // Highly active lifestyle
)

// SurveyData represents the data collected in a survey.
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
