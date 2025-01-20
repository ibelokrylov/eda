package main

import (
	"eda/app/config"
	"eda/app/handlers"
	"eda/app/middleware"
	"eda/app/service"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

const Version = "1.0.0"

func main() {
	config.Init()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowOriginsFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost:") || origin == "http://localhost" {
				return true
			}

			allowedOrigins := []string{"https://example.com", "https://another.example.com"}
			for _, o := range allowedOrigins {
				if o == origin {
					return true
				}
			}
			return false
		},
	}))
	app.Use(func(c *fiber.Ctx) error {
		return middleware.LoggerMiddleware(c, config.Logger)
	})

	port := config.GetEnvVariable("APP_PORT")
	// TODO: add clear cookie else invalid cookie
	api := app.Group("/api")
	api.Use(healthcheck.New())

	v1 := api.Group("/v1")

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

	go service.ParseProduct()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	log.Fatal(app.Listen(port))

	for {
		select {
		case <-ticker.C:
			go service.ParseProduct()
		}
	}
}
