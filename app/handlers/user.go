package handlers

import (
	"safechron/api/app/config"
	"safechron/api/app/entities"
	"safechron/api/app/helpers"
	"safechron/api/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.CreateUser)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}

	if validationErr := helpers.ValidateStruct(user); validationErr != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), validationErr.Error()))
	}

	new_user, err := service.CreateUser(*user)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	createSessionErr := config.SetSessionKey(c, "user", config.UserSession{ID: new_user.ID})
	if createSessionErr != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), createSessionErr.Error()))
	}

	// TODO: Send email to user code
	// code, err := service.GenerateCode(new_user.ID, "registration")
	// if err != nil {
	// 	return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	// }
	// err_email := service.SendEmail([]string{new_user.Username}, service.EmailData{Subject: "Account registration"}, "registration", map[string]string{"code": code.Code})
	// if err_email != nil {
	// 	return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err_email.Error()))
	// }

	return c.JSON(config.BaseResult(config.GetStatus("OK"), new_user))
}

type GetUserOptions struct {
	IncludeDeleted bool
}

func GetUserById(user_id uuid.UUID) (entities.User, error) {
	u, err := service.GetUserById(user_id)
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetUserProfile(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	user_data, err := GetUserById(user.ID)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	return c.JSON(config.BaseResult(config.GetStatus("OK"), user_data))
}

func UserLogout(c *fiber.Ctx) error {
	config.DeleteSessionKey(c, "user")
	return c.JSON(config.BaseResult(config.GetStatus("OK"), "User logged out"))
}

func GetCodeRegistration(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), err.Error()))
	}
	return service.GetUserRegistrationNewOrOldCode(user.ID)
}
