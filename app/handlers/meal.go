package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"eda/app/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateMeal(c *fiber.Ctx) error {
	u, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	cm := new(entities.CreateMeal)

	if err := c.BodyParser(cm); err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	err = helpers.ValidateStruct(cm)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	d, err := time.Parse("2006-01-02", cm.Day)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, "cerate-meal-hn@Date format error - YYYY-MM-DD"))
	}

	meal, err := service.CreateMeal(cm, u.ID, d)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}
	if meal.ID == 0 {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, "cerate-meal-hn@Meal not created"))
	}
	return c.JSON(config.BaseResult(config.GetStatus("OK"), meal))
}

func GetMealByUserId(c *fiber.Ctx) error {
	u, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	date, err := time.Parse("2006-01-02", c.Query("date"))
	if err != nil {
		date = time.Now().Truncate(24 * time.Hour)
	}

	meal, err := service.GetMealByUserId(u.ID, date)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}
	return c.JSON(config.BaseResult(config.GetStatus("OK"), meal))
}
