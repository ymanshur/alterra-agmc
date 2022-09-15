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

func GetUser(userId int) (interface{}, error) {
	user := new(model.User)
	if err := config.DB.First(user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(userId int, user *model.User) (interface{}, error) {
	updatedUser := new(model.User)
	if err := config.DB.Model(updatedUser).Where("id = ?", userId).Updates(user).Error; err != nil {
		return err, nil
	}
	return updatedUser, nil
}

func DeleteUser(userId int) error {
	if err := config.DB.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUser() (interface{}, error) {
	users := new([]model.User)
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
