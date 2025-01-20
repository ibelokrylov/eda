package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"eda/app/service"

	"github.com/gofiber/fiber/v2"
)

func FoodCreate(c *fiber.Ctx) error {
	u, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}
	f := new(entities.CreateFood)
	if err := c.BodyParser(f); err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	err = helpers.ValidateStruct(f)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	cf, err := service.CreateFood(*f, u.ID)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	if cf.ID == 0 {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, "cerate-food-hn@Food not created"))
	}
	return c.JSON(config.BaseResult(config.GetStatus("OK"), cf))
}
