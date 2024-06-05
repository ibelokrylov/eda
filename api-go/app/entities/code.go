package entities

import (
	"time"

	"github.com/google/uuid"
)

type Code struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UserId    uuid.UUID `json:"user_id" validate:"required,uuid4" gorm:"not null" sql:"type:uuid"`
	Type      string    `json:"type" validate:"required,oneof=reset_password,change_email,registration,confirm_auth"`
	Code      string    `json:"code" validate:"required,min=6,max=6"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
}
