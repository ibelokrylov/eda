package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"strconv"
	"time"
)

func GenerateCode(user_id int64, code_type entities.CodeType) (entities.Code, error) {
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

func GetUserCodeByType(user_id int64, code_type entities.CodeType) (entities.Code, error) {
	var code entities.Code
	now := time.Now()
	c := config.Db.Where("id = ? AND type = ? AND is_used = false AND created_at >= ?", user_id, code_type, now.Add(-30*time.Minute)).Find(&code)

	if c.Error != nil {
		return GenerateCode(user_id, code_type)
	}
	return code, nil
}
