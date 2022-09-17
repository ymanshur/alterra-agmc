package repository

import (
	"day-4/go-restful/lib/auth"
	"day-4/go-restful/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Get(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := r.DB.Omit("token", "password").Where(model.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(userId uint, user *model.User) (*model.User, error) {
	if err := r.DB.Model(&model.User{ID: userId}).Updates(user).Error; err != nil {
		return nil, err
	}

	updatedUser := new(model.User)
	if err := r.DB.Omit("token", "password").Find(&updatedUser, userId).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (r *UserRepository) Delete(userId uint) error {
	if err := r.DB.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Omit("token", "password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Login(user *model.User) (*model.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("required both email and password")
	}

	if err := r.DB.Where(model.User{Email: user.Email, Password: user.Password}).Take(&user).Error; err != nil {
		return nil, errors.New("these credentials do not match our records")
	}

	token, err := auth.CreateJwt(user)
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}
