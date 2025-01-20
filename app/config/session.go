package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session *session.Store

type UserSession struct {
	ID int64 `json:"id" validate:"required"`
}

func InitSession() {
	day, err := strconv.Atoi(GetEnvVariable("SESSION_EXPIRATION"))
	if err != nil {
		day = 1
	}
	day = day * 24

	Session = session.New(session.Config{
		Expiration:     time.Duration(day) * time.Hour,
		Storage:        GetRedis(),
		CookieSameSite: fiber.CookieSameSiteNoneMode,
	})
}

func ParseUserSession(c *fiber.Ctx) (UserSession, error) {
	sess, err := Session.Get(c)
	if err != nil {
		return UserSession{}, err
	}

	user := sess.Get("user")
	if user == nil {
		return UserSession{}, fmt.Errorf("user not found in session")
	}

	userSessionPointer, ok := user.(*UserSession)
	if !ok {
		return UserSession{}, fmt.Errorf("user session is not of type *config.UserSession")
	}

	if userSessionPointer == nil {
		return UserSession{}, fmt.Errorf("user session pointer is nil")
	}

	return *userSessionPointer, nil
}

func GetSession() *session.Store {
	return Session
}

func GetSessionKey(c *fiber.Ctx, key string) interface{} {
	store := GetSession()
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	value := sess.Get(key)
	return value
}

func SetSessionKey(c *fiber.Ctx, key string, value interface{}) error {
	store := GetSession()
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	sess.Set(key, value)
	err = sess.Save()
	if err != nil {
		return err
	}
	return nil
}

func UpdateSessionKey(c *fiber.Ctx, key string, value interface{}) error {
	store := GetSession()
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	sess.Set(key, value)
	err = sess.Save()
	if err != nil {
		return err
	}
	return nil
}

func DeleteSessionKey(c *fiber.Ctx, key string) {
	store := GetSession()
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Delete(key)
	if err := sess.Destroy(); err != nil {
		panic(err)
	}
	err = sess.Save()
	if err != nil {
		panic(err)
	}
}
