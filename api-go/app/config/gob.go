package config

import (
	"encoding/gob"
	"safechron/api/app/entities"

	"github.com/google/uuid"
)

func InitGob() {
	gob.Register(uuid.UUID{})
	gob.Register(&UserSession{})
	gob.Register(&entities.LoginUser{})
}
