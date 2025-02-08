package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index,default:NULL"`
	Username     string         `json:"username" validate:"required,email"`
	EmailConfirm bool           `json:"email_confirm" gorm:"default:false"`
	Password     string         `json:"-" validate:"required,min=8"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	FirstName    string         `json:"first_name" validate:"required"`
	LastName     string         `json:"last_name" validate:"required"`
	Survey       UserSurvey     `json:"-" gorm:"foreignKey:UserID"`
	Meal         Meal           `json:"-" gorm:"foreignKey:UserID"`
	BzuNorm      []UserBzuNorm  `json:"-" gorm:"foreignKey:UserID"`
}

type CreateUser struct {
	Username        string `json:"username" validate:"required,min=3"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"passwordConfirm" validate:"required,eqfield=Password"`
	FirstName       string `json:"firstName" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
}

type UserBzuNorm struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Max       float64   `json:"max"`
	Day       time.Time `json:"day" validate:"required" gorm:"default:CURRENT_TIMESTAMP;index"`
	Carb      float64   `json:"carbs"`
	Fat       float64   `json:"fat"`
	Protein   float64   `json:"protein"`
}

type UserBzuCalculate struct {
	Max     float64
	Carb    float64
	Fat     float64
	Protein float64
}

type UserBzuNormResponse struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Max       float64   `json:"max"`
	Day       time.Time `json:"day" validate:"required" gorm:"default:CURRENT_TIMESTAMP;index"`
	Current   float64   `json:"current"`
	Carb      float64   `json:"carbs"`
	Fat       float64   `json:"fat"`
	Protein   float64   `json:"protein"`
}

type CreateUserBzuNorm struct {
	Day string `json:"day" validate:"required"`
}
