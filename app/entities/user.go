package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username      string         `json:"username" validate:"required,email"`
	Email_confirm bool           `json:"email_confirm" gorm:"default:false"`
	Password      string         `json:"-" validate:"required,min=8"`
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	FirstName     string         `json:"first_name" validate:"required"`
	LastName      string         `json:"last_name" validate:"required"`
	Survey        UserSurvey     `json:"-" gorm:"foreignKey:UserID"`
	Meal          Meal           `json:"-" gorm:"foreignKey:UserID"`
}

type CreateUser struct {
	Username        string `json:"username" validate:"required,min=3"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"passwordConfirm" validate:"required,eqfield=Password"`
	FirstName       string `json:"firstName" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
}

func (u *User) UserDelete(tx *gorm.DB) (err error) {
	var zeroTime time.Time
	if u.DeletedAt.Time == zeroTime {
		u.DeletedAt.Time = time.Now()
	}
	return
}
