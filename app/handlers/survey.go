package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/service"

	"github.com/gofiber/fiber/v2"
)

func CreateSurvey(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)

	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	body_survey := new(entities.SurveyData)
	if err := c.BodyParser(body_survey); err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	survey, err := service.CreateSurvey(user.ID, *body_survey)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}

	return c.JSON(config.BaseResult(config.GetStatus("OK"), survey))
}

func GetSurveyByUserId(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	survey, err := service.GetSurveyByUserId(user.ID)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	return c.JSON(config.BaseResult(config.GetStatus("OK"), survey))
}
