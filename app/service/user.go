package service

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/helpers"
	"errors"

	"github.com/google/uuid"
)

func CreateUser(user entities.CreateUser) (*entities.User, error) {
	find_user, _ := GetUserByUsername(user.Username)
	if find_user.ID != uuid.Nil {
		return nil, errors.New("user not created")
	}

	new_user := new(entities.User)
	new_user.Username = user.Username
	new_user.IsActive = false
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName
	hash_password, _ := helpers.HashPassword(user.Password)
	new_user.Password = hash_password

	u := config.Db.Create(&new_user)
	if u.Error != nil {
		return nil, u.Error
	}
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

func GetUserById(user_id uuid.UUID) (entities.User, error) {
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

func GetUserRegistrationNewOrOldCode(user_id uuid.UUID) error {
	code, err := GetUserCodeByType(user_id, entities.CODE_REGISTRATION)
	if err != nil {
		return err
	}
	user, _ := GetUserById(user_id)
	SendEmailCodeConfirmRegistration(user.Username, code.Code)
	return nil
}
