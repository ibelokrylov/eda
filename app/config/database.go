package config

import (
	"eda/app/entities"
	"fmt"

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

	Db.AutoMigrate(
		&entities.User{},
		&entities.Code{},
		&entities.UserSurvey{},
		&entities.ProductCategory{},
		&entities.Food{},
		&entities.Product{},
		&entities.ProductParesed{},
		&entities.UserBzuNorm{},
		&entities.Meal{},
	)
	return nil
}

func UpdateTSV(db, name, lang, content string, id int64) error {
	sql := fmt.Sprintf(`
        UPDATE %s
        SET %s = to_tsvector(%s, ?)
        WHERE id = ?
    `, db, name, lang)

	return Db.Exec(sql, content, id).Error
}
