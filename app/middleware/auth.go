package middleware

import (
	"eda/app/config"
	"eda/app/service"

	"github.com/gofiber/fiber/v2"
)

func IsAuthRequired(c *fiber.Ctx) error {
	user, err_parse_session := config.ParseUserSession(c)
	if err_parse_session != nil {
		return c.JSON(config.BaseResult(config.GetStatus("NOT_AUTH"), "Unauthorized"))
	}

	if user.ID == 0 {
		return c.JSON(config.BaseResult(config.GetStatus("NOT_AUTH"), "Unauthorized"))
	}
	_, err := service.GetUserById(user.ID)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("NOT_AUTH"), "Unauthorized"))
	}

	return c.Next()
}

func IsNeedDontAuth(c *fiber.Ctx) error {
	err := IsAuthRequired(c)
	if err == nil {
		return c.Next()
	}
	return c.JSON(config.BaseResult(config.GetStatus("FAIL"), "Authenticated"))
}
