package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserSurvey struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId    uuid.UUID  `json:"user_id" validate:"required,uuid4"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime:true"`
	Data      SurveyData `json:"data" validate:"required" gorm:"JSON"`
}

type Target string

const (
	LOSE_WEIGHT     Target = "LOSE_WEIGHT"     //Потеря веса
	MAINTAIN_WEIGHT Target = "MAINTAIN_WEIGHT" //Поддержание веса
	GAIN_WEIGHT     Target = "GAIN_WEIGHT"     //Приобретение веса
)

type Activity string

const (
	SEDENTARY_LIFESTYLE     Activity = "SEDENTARY_LIFESTYLE"     //Сидячая жизнь
	MODERATE_LIFESTYLE      Activity = "MODERATE_LIFESTYLE"      //Умеренная жизнь
	ACTIVE_LIFESTYLE        Activity = "ACTIVE_LIFESTYLE"        //Активная жизнь
	HIGHLY_ACTIVE_LIFESTYLE Activity = "HIGHLY_ACTIVE_LIFESTYLE" //Высокая активная жизнь
)

type SurveyData struct {
	Gender   string   `json:"gender" validate:"required,eq=MALE|eq=FEMALE"`
	Target   Target   `json:"target" validate:"required,eq=LOSE_WEIGHT|eq=MAINTAIN_WEIGHT|eq=GAIN_WEIGHT"`
	Age      int      `json:"age" validate:"required"`
	Growth   int      `json:"growth" validate:"required"`
	Bithday  string   `json:"birthday" validate:"required,date"`
	Activity Activity `json:"activity" validate:"required,eq=SEDENTARY_LIFESTYLE|eq=MODERATE_LIFESTYLE|eq=ACTIVE_LIFESTYLE|eq=HIGHLY_ACTIVE_LIFESTYLE"`
	Weight   int      `json:"weight" validate:"required"`
}
