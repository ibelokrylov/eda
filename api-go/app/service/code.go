package service

import (
	"safechron/api/app/config"
	"safechron/api/app/entities"
	"safechron/api/app/helpers"
	"strconv"

	"github.com/google/uuid"
)

func GenerateCode(user_id uuid.UUID, code_type string) (entities.Code, error) {
	var code entities.Code

	code.UserId = user_id
	code.Type = code_type
	code.Code = strconv.Itoa(helpers.GenerateCode())

	u := config.Db.Create(&code)
	if u.Error != nil {
		return entities.Code{}, u.Error
	}
	return code, nil
}
