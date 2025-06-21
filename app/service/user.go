package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"errors"
	"gorm.io/gorm"
	"time"
)

func CreateUser(user entities.CreateUser) (*entities.User, error) {
	// Check if a user with the given username already exists
	fu, _ := GetUserByUsername(user.Username)
	if fu.ID != 0 {
		return nil, errors.New("cretae-user@user-not-created")
	}

	// Create a new user instance
	nu := new(entities.User)
	nu.Username = user.Username
	nu.IsActive = false
	nu.FirstName = user.FirstName
	nu.LastName = user.LastName

	// Hash the user's password
	hp, _ := helpers.HashPassword(user.Password)
	nu.Password = hp

	// Save the new user to the database
	u := config.Db.Create(&nu)
	if u.Error != nil {
		return nil, u.Error
	}

	// Uncomment and implement the following lines if email confirmation is required
	// code, err := GenerateCode(nu.ID, entities.CODE_REGISTRATION)
	// if err != nil {
	// 	return nu, err
	// }
	// err = SendEmailCodeConfirmRegistration(nu.Username, code.Code)
	// if err != nil {
	// 	return nu, err
	// }

	return nu, nil
}

func GetUserById(userId int64) (entities.User, error) {
	var user entities.User
	u := config.Db.Unscoped().First(
		&user,
		"id = ?",
		userId,
	)
	if u.Error != nil {
		return entities.User{}, u.Error
	}
	return user, nil
}

func GetUserByUsername(username string) (entities.User, error) {
	var user entities.User
	u := config.Db.Unscoped().First(
		&user,
		"username = ?",
		username,
	)
	if u.Error != nil {
		return entities.User{}, u.Error
	}
	return user, nil
}

func GetUserRegistrationNewOrOldCode(userId int64) error {
	code, err := GetUserCodeByType(
		userId,
		entities.CODE_REGISTRATION,
	)
	if err != nil {
		return err
	}
	user, _ := GetUserById(userId)
	err = SendEmailCodeConfirmRegistration(
		user.Username,
		code.Code,
	)
	if err != nil {
		return err
	}
	return nil
}

func RegenerateUserBzuNorm(id int64) error {
	date := time.Now().Truncate(24 * time.Hour)

	cb, err := CalculatedUserBzu(id)
	if err != nil {
		return err
	}

	if err := config.Db.
		Where("user_id = ?", id).
		Where("day = ?", date).
		First(&entities.UserBzuNorm{}).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			bzu := &entities.UserBzuNorm{
				UserID:  id,
				Day:     date,
				Max:     cb.Max,
				Fat:     cb.Fat,
				Protein: cb.Protein,
				Carb:    cb.Carb,
			}

			config.Db.Create(&bzu)
			return nil
		}
		return err
	}

	err = config.Db.
		Where("user_id = ?", id).
		Where("day = ?", date).
		Updates(entities.UserBzuNorm{
			Max:     cb.Max,
			Fat:     cb.Fat,
			Protein: cb.Protein,
			Carb:    cb.Carb,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func GenerateOrReadBzu(id int64, date time.Time) (entities.UserBzuNormResponse, error) {
	bzu := new(entities.UserBzuNorm)
	res := new(entities.UserBzuNormResponse)
	if err := config.Db.Where(
		"user_id = ?",
		id,
	).Where(
		"day = ?",
		date,
	).First(bzu).Error; err != nil {
		if err.Error() == "record not found" {
			// generate bzu
			cb, err := CalculatedUserBzu(id)
			if err != nil {
				return entities.UserBzuNormResponse{}, err
			}

			bzu = &entities.UserBzuNorm{
				UserID:  id,
				Day:     date,
				Max:     cb.Max,
				Fat:     cb.Fat,
				Protein: cb.Protein,
				Carb:    cb.Carb,
			}
			config.Db.Create(&bzu)
		} else {
			return entities.UserBzuNormResponse{}, err
		}
	}
	res.ID = bzu.ID
	res.UserID = bzu.UserID
	res.CreatedAt = bzu.CreatedAt
	res.UpdatedAt = bzu.UpdatedAt
	res.Max = bzu.Max
	res.Day = bzu.Day

	m, err := GetMealByUserId(id, date)
	if err != nil {
		return entities.UserBzuNormResponse{}, err
	}

	uBzu, err := CalculatedUserBzu(id)

	if len(m) != 0 {
		for _, meal := range m {
			res.Current += meal.Info.Calories
		}
	}

	res.Fat = uBzu.Fat
	res.Protein = uBzu.Protein
	res.Carb = uBzu.Carb

	return *res, nil
}

func CalculatedUserBzu(id int64) (*entities.UserBzuCalculate, error) {
	u := new(entities.User)

	if err := config.Db.Preload("Survey").Where(
		"id = ?",
		id,
	).First(u).Error; err != nil {
		return nil, err
	}
	r := new(entities.UserBzuCalculate)

	var calc float64 = 0
	age := CalculateUserAgeByBirthdayDate(u.Survey.Data.Birthday)

	if u.Survey.Data.Gender == "" {
		return nil, errors.New("CalculatedUserBzu@Survey-required")
	}
	if u.Survey.Data.Gender == "MALE" {
		calc = 66.5 + (13.75 * float64(u.Survey.Data.Weight)) + (5.003 * float64(u.Survey.Data.Growth)) - (6.755 * float64(age))
	} else {
		calc = 655.1 + (9.563 * float64(u.Survey.Data.Weight)) + (1.850 * float64(u.Survey.Data.Growth)) - (4.676 * float64(age))
	}

	cp := 0.227
	cf := 0.299
	cc := 0.474
	switch u.Survey.Data.Activity {
	case "SEDENTARY_LIFESTYLE":
		calc = calc * 1.2
	case "MODERATE_LIFESTYLE":
		calc = calc * 1.375
	case "ACTIVE_LIFESTYLE":
		calc = calc * 1.55
	case "HIGHLY_ACTIVE_LIFESTYLE":
		calc = calc * 1.725
	}

	switch u.Survey.Data.Target {
	case "LOSE_WEIGHT":
		calc = calc - 400
		cp = 0.301
		cf = 0.338
		cc = 0.359
	case "GAIN_WEIGHT":
		calc = calc + 400
		cp = 0.185
		cf = 0.293
		cc = 0.522
	default:
	}

	r.Max = calc
	r.Carb = calc * cc / 4
	r.Fat = calc * cf / 9
	r.Protein = calc * cp / 4

	return r, nil
}

func CalculateUserAgeByBirthdayDate(bithday time.Time) int {
	age := time.Now().Year() - bithday.Year()
	if bithday.Month() > time.Now().Month() {
		age--
	}
	if bithday.Month() == time.Now().Month() && bithday.Day() > time.Now().Day() {
		age--
	}
	return age
}
