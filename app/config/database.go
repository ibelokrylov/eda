package config

import (
	"eda/app/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() error {
	var err error
	var DATABASE_URL string = GetEnvVariable("DATABASE_URL")

	Db, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&entities.User{}, &entities.Code{}, &entities.UserSurvey{})
	return nil
}
