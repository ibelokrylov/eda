package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
)

func CreateSurvey(userID int64, survey entities.SurveyData) (entities.UserSurvey, error) {
	findSurvey, _ := GetSurveyByUserId(userID)

	if findSurvey.ID != 0 {
		if err := helpers.ValidateStruct(&survey); err != nil {
			return entities.UserSurvey{}, err
		}
		findSurvey.Data = survey
		if err := config.Db.Save(&findSurvey).Error; err != nil {
			return entities.UserSurvey{}, err
		}

		err := RegenerateUserBzuNorm(userID)
		if err != nil {
			return entities.UserSurvey{}, err
		}
		return findSurvey, nil
	}

	newSurvey := new(entities.UserSurvey)

	newSurvey.UserID = userID
	newSurvey.Data = survey

	er := config.Db.Create(&newSurvey)

	if er.Error != nil {
		return entities.UserSurvey{}, er.Error
	}

	err := RegenerateUserBzuNorm(userID)
	if err != nil {
		return entities.UserSurvey{}, err
	}

	return *newSurvey, nil
}

func GetSurveyByUserId(userID int64) (entities.UserSurvey, error) {
	var survey entities.UserSurvey
	u := config.Db.First(&survey, "user_id = ?", userID)
	if u.Error != nil {
		return entities.UserSurvey{}, u.Error
	}
	return survey, nil
}
