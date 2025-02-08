package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"eda/app/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.CreateUser)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}

	if err := helpers.ValidateStruct(user); err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}

	new_user, err := service.CreateUser(*user)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}
	createSessionErr := config.SetSessionKey(
		c,
		"user",
		config.UserSession{ID: new_user.ID},
	)
	if createSessionErr != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				createSessionErr.Error(),
			),
		)
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

	return c.JSON(
		config.BaseResult(
			config.GetStatus("OK"),
			new_user,
		),
	)
}

type GetUserOptions struct {
	IncludeDeleted bool
}

func GetUserById(user_id int64) (entities.User, error) {
	u, err := service.GetUserById(user_id)
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetUserProfile(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}
	user_data, err := GetUserById(user.ID)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}
	return c.JSON(
		config.BaseResult(
			config.GetStatus("OK"),
			user_data,
		),
	)
}

func UserLogout(c *fiber.Ctx) error {
	config.DeleteSessionKey(
		c,
		"user",
	)
	return c.JSON(
		config.BaseResult(
			config.GetStatus("OK"),
			"User logged out",
		),
	)
}

func GetCodeRegistration(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}
	return service.GetUserRegistrationNewOrOldCode(user.ID)
}

func GetUserBzu(c *fiber.Ctx) error {
	user, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}

	date, err := time.Parse(
		"2006-01-02",
		c.Query("date"),
	)
	if err != nil {
		date = time.Now().Truncate(24 * time.Hour)
	}
	bzu, err := service.GenerateOrReadBzu(
		user.ID,
		date,
	)
	if err != nil {
		return c.JSON(
			config.BaseResult(
				config.GetStatus("FAIL"),
				nil,
				err.Error(),
			),
		)
	}
	return c.JSON(
		config.BaseResult(
			config.GetStatus("OK"),
			bzu,
		),
	)
}
