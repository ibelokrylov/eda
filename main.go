package main

import (
	"eda/app/config"
	"eda/app/middleware"
	"eda/app/routes"
	"eda/app/service"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	routes.Setup(app)
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
