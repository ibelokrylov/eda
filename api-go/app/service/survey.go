package service

import (
	"errors"
	"safechron/api/app/config"
	"safechron/api/app/entities"

	"github.com/google/uuid"
)

func CreateSurvey(userUUID uuid.UUID, survey entities.SurveyData) (entities.UserSurvey, error) {
	find_survey, _ := GetSurveyByUserId(userUUID)

	if find_survey.ID != uuid.Nil {
		return entities.UserSurvey{}, errors.New("user already have a survey")
	}

	new_survey := new(entities.UserSurvey)

	new_survey.UserId = userUUID
	new_survey.Data = survey

	er := config.Db.Create(&new_survey)

	if er.Error != nil {
		return entities.UserSurvey{}, er.Error
	}

	return *new_survey, nil
}

func GetSurveyByUserId(userUUID uuid.UUID) (entities.UserSurvey, error) {
	var survey entities.UserSurvey
	u := config.Db.Unscoped().First(&survey, "user_id = ?", userUUID)
	if u.Error != nil {
		return entities.UserSurvey{}, u.Error
	}
	return survey, nil
}
