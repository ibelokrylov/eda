package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"eda/app/service"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	user := new(entities.LoginUser)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}

	if validationErr := helpers.ValidateStruct(user); validationErr != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), validationErr.Error()))
	}

	find_user, find_err := service.GetUserByUsername(user.Username)
	if find_err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), find_err.Error()))
	}

	if !helpers.CheckPasswordHash(user.Password, find_user.Password) {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), "Incorrect password"))
	}

	createSessionErr := config.SetSessionKey(c, "user", config.UserSession{ID: find_user.ID})
	if createSessionErr != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), createSessionErr.Error()))
	}

	// TODO: Send email to user

	return c.JSON(config.BaseResult(config.GetStatus("OK"), find_user))
}
