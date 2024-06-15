package main

import (
	"log"
	"safechron/api/app/config"
	"safechron/api/app/handlers"
	"safechron/api/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

const Version = "1.0.0"

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	config.Init()

	port := config.GetEnvVariable("APP_PORT")
	// TODO: add clear cookie else invalid cookie
	api := app.Group("/api")
	api.Use(healthcheck.New())

	v1 := api.Group("/v1")

	user := v1.Group("/user")
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

	log.Fatal(app.Listen(port))
}
