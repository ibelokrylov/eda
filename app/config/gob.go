package config

import (
	"eda/app/entities"
	"encoding/gob"

	"github.com/google/uuid"
)

func InitGob() {
	gob.Register(uuid.UUID{})
	gob.Register(&UserSession{})
	gob.Register(&entities.LoginUser{})
}
