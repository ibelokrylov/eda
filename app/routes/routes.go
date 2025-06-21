package routes

import (
	"eda/app/handlers"
	"eda/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Use(healthcheck.New())

	v1 := api.Group("/v1")

	test := v1.Group("/test")
	test.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})

	user := v1.Group("/user")
	user.Post("/meal", middleware.IsAuthRequired, handlers.CreateMeal)
	user.Get("/meal", middleware.IsAuthRequired, handlers.GetMealByUserId)
	user.Get("/bzu", middleware.IsAuthRequired, handlers.GetUserBzu)

	registerUser := user.Group("/register", middleware.IsNeedDontAuth)
	registerUser.Post("/", handlers.CreateUser)

	userProfile := user.Group("/profile", middleware.IsAuthRequired)
	userProfile.Get("/", handlers.GetUserProfile)

	userLogout := user.Group("/logout", middleware.IsAuthRequired)
	userLogout.Post("/", handlers.UserLogout)

	userAuth := user.Group("/auth", middleware.IsNeedDontAuth)
	userAuth.Post("/", handlers.Authenticated)

	userSurvey := user.Group("/survey", middleware.IsAuthRequired)
	userSurvey.Post("/", handlers.CreateSurvey)
	userSurvey.Get("/", handlers.GetSurveyByUserId)

	emailUser := user.Group("/email", middleware.IsAuthRequired)
	emailUser.Get("/registration_code", handlers.GetCodeRegistration)

	search := v1.Group("/search", middleware.IsAuthRequired)
	search.Get("/product", handlers.FindFoodAndProductAndCategories)

	food := v1.Group("/food", middleware.IsAuthRequired)
	food.Post("/", handlers.FoodCreate)
}
