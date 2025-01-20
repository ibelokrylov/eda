package middleware

import (
	"eda/app/config"
	"eda/app/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(c *fiber.Ctx, logger *logrus.Logger) error {
	traceID := c.Cookies("X-Request-Id")
	userSession, _ := config.ParseUserSession(c)
	sessionID := "not_authenticated"

	if userSession.ID != 0 {
		sessionID = strconv.FormatInt(userSession.ID, 10)
	}

	if traceID == "" {
		newTraceID := strconv.FormatInt(helpers.GenerateRandomInt64(), 10)
		c.Cookie(&fiber.Cookie{
			Name:  "X-Request-Id",
			Value: newTraceID,
		})
		traceID = newTraceID
	}

	entry := logger.WithFields(logrus.Fields{
		"method":     c.Method(),
		"path":       c.Path(),
		"ip":         c.IP(),
		"user_id":    sessionID,
		"status":     c.Response().StatusCode(),
		"data":       string(c.Body()),
		"user_agent": c.Get("User-Agent"),
	})
	entry.Info(traceID)
	return c.Next()
}
