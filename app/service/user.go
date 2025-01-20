package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"errors"
	"time"
)

func CreateUser(user entities.CreateUser) (*entities.User, error) {
	// Check if a user with the given username already exists
	find_user, _ := GetUserByUsername(user.Username)
	if find_user.ID != 0 {
		return nil, errors.New("cretae-user@user-not-created")
	}

	// Create a new user instance
	new_user := new(entities.User)
	new_user.Username = user.Username
	new_user.IsActive = false
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName

	// Hash the user's password
	hash_password, _ := helpers.HashPassword(user.Password)
	new_user.Password = hash_password

	// Save the new user to the database
	u := config.Db.Create(&new_user)
	if u.Error != nil {
		return nil, u.Error
	}

	// Uncomment and implement the following lines if email confirmation is required
	// code, err := GenerateCode(new_user.ID, entities.CODE_REGISTRATION)
	// if err != nil {
	// 	return new_user, err
	// }
	// err = SendEmailCodeConfirmRegistration(new_user.Username, code.Code)
	// if err != nil {
	// 	return new_user, err
	// }

	return new_user, nil
}

func GetUserById(user_id int64) (entities.User, error) {
	var user entities.User
	u := config.Db.Unscoped().First(&user, "id = ?", user_id)
	if u.Error != nil {
		return entities.User{}, u.Error
	}
	return user, nil
}

func GetUserByUsername(username string) (entities.User, error) {
	var user entities.User
	u := config.Db.Unscoped().First(&user, "username = ?", username)
	if u.Error != nil {
		return entities.User{}, u.Error
	}
	return user, nil
}

func UpdateUser(user entities.User) (entities.User, error) {
	u := config.Db.Save(&user)
	if u.Error != nil {
		return entities.User{}, u.Error
	}
	return user, nil
}

func GetUserRegistrationNewOrOldCode(user_id int64) error {
	code, err := GetUserCodeByType(user_id, entities.CODE_REGISTRATION)
	if err != nil {
		return err
	}
	user, _ := GetUserById(user_id)
	SendEmailCodeConfirmRegistration(user.Username, code.Code)
	return nil
}

func GenerateOrReadBzu(id int64, date time.Time) (entities.UserBzuNormResponse, error) {
	bzu := new(entities.UserBzuNorm)
	res := new(entities.UserBzuNormResponse)
	if err := config.Db.Where("user_id = ?", id).Where("day = ?", date).First(bzu).Error; err != nil {
		if err.Error() == "record not found" {
			// generate bzu
			cb, err := CalculatedUserBzu(id)
			if err != nil {
				return entities.UserBzuNormResponse{}, err
			}

			bzu = &entities.UserBzuNorm{
				UserID: id,
				Day:    date,
				Max:    cb,
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

	m := new(entities.Meal)
	ferr := config.Db.Where("day = ?", date).Where("user_id = ?", id).First(m).Error
	if ferr == nil {
		c := 0.0

		if len(m.MealFoods) != 0 {
			for _, i := range m.MealFoods {
				cf := i.Weight / 100
				c += i.Info.Calories * cf
			}
		}

		res.Current = c
	}

	return *res, nil
}

func CalculatedUserBzu(id int64) (float64, error) {
	u := new(entities.User)

	if err := config.Db.Preload("Survey").Where("id = ?", id).First(u).Error; err != nil {
		return 0, err
	}

	var calc float64 = 0
	age := CalculateUserAgeByBirthdayDate(u.Survey.Data.Birthday)

	if u.Survey.Data.Gender == "" {
		return 0, errors.New("CalculatedUserBzu@Survey-required")
	}
	if u.Survey.Data.Gender == "MALE" {
		calc = 66.5 + (13.75 * float64(u.Survey.Data.Weight)) + (5.003 * float64(u.Survey.Data.Growth)) - (6.755 * float64(age))
	} else {
		calc = 655.1 + (9.563 * float64(u.Survey.Data.Weight)) + (1.850 * float64(u.Survey.Data.Growth)) - (4.676 * float64(age))
	}

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
	case "GAIN_WEIGHT":
		calc = calc + 400
	default:
	}

	return calc, nil
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
