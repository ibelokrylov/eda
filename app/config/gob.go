package config

import (
	"eda/app/entities"
	"encoding/gob"
)

func InitGob() {
	gob.Register(&UserSession{})
	gob.Register(&entities.LoginUser{})
}
