package middleware

import (
	"safechron/api/app/config"
	"safechron/api/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func IsAuthRequired(c *fiber.Ctx) error {
	user, err_parse_session := config.ParseUserSession(c)
	if err_parse_session != nil {
		return c.JSON(config.BaseResult(config.GetStatus("NOT_AUTH"), "Unauthorized"))
	}

	user_id, parse_err_uuid := uuid.Parse(user.ID.String())
	if parse_err_uuid != nil {
		return c.JSON(config.BaseResult(config.GetStatus("NOT_AUTH"), "Unauthorized"))
	}
	_, err := service.GetUserById(user_id)
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
