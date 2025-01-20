package entities

import (
	"time"
)

type CodeType string

const (
	CODE_RESET_PASSWORD CodeType = "reset_password"
	CODE_CHANGE_EMAIL   CodeType = "change_email"
	CODE_REGISTRATION   CodeType = "registration"
	CODE_CONFIRM_AUTH   CodeType = "confirm_auth"
)

type Code struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UserId    int64     `json:"user_id" validate:"required" gorm:"not null"`
	Type      CodeType  `json:"type" validate:"required,eq=reset_password|eq=change_email|eq=registration|eq=confirm_auth"`
	Code      string    `json:"code" validate:"required,min=6,max=6"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
}
