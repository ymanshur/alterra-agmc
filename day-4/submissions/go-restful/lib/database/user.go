package database

import (
	"go-restful/config"
	"go-restful/model"
)

func CreateUser(user *model.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(userId uint) (interface{}, error) {
	user := new(model.User)
	if err := config.DB.Omit("token", "password").Where(model.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(userId uint, user *model.User) (interface{}, error) {
	if err := config.DB.Model(&model.User{ID: userId}).Updates(user).Error; err != nil {
		return err, nil
	}

	updatedUser := new(model.User)
	if err := config.DB.Omit("token", "password").Find(&updatedUser, userId).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func DeleteUser(userId uint) error {
	if err := config.DB.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUser() (interface{}, error) {
	users := new([]model.User)
	if err := config.DB.Omit("token", "password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
