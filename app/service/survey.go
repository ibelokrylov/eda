package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"

	"github.com/google/uuid"
)

func CreateSurvey(userUUID uuid.UUID, survey entities.SurveyData) (entities.UserSurvey, error) {
	find_survey, _ := GetSurveyByUserId(userUUID)

	if find_survey.ID != uuid.Nil {
		if err := helpers.ValidateStruct(&survey); err != nil {
			return entities.UserSurvey{}, err
		}
		find_survey.Data = survey
		if err := config.Db.Save(&find_survey).Error; err != nil {
			return entities.UserSurvey{}, err
		}
		return find_survey, nil
	}

	new_survey := new(entities.UserSurvey)

	new_survey.UserID = userUUID
	new_survey.Data = survey

	er := config.Db.Create(&new_survey)

	if er.Error != nil {
		return entities.UserSurvey{}, er.Error
	}

	return *new_survey, nil
}

func GetSurveyByUserId(userUUID uuid.UUID) (entities.UserSurvey, error) {
	var survey entities.UserSurvey
	u := config.Db.First(&survey, "user_id = ?", userUUID)
	if u.Error != nil {
		return entities.UserSurvey{}, u.Error
	}
	return survey, nil
}
