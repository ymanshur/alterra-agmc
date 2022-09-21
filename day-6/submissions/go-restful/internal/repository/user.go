package repository

import (
	"day-6/go-restful/internal/model"
	"day-6/go-restful/internal/pkg/util"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u *User) Save(user *model.User) (*model.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) FindById(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("token", "password").Where(model.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdateById(userId uint, user *model.User) (*model.User, error) {
	if err := u.db.Model(&model.User{ID: userId}).Updates(&user).Error; err != nil {
		return nil, err
	}

	updatedUser := new(model.User)
	if err := u.db.Omit("token", "password").Find(&updatedUser, userId).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *User) DeleteById(userId uint) error {
	if err := u.db.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Omit("token", "password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) FindByEmailPassword(user *model.User) (*model.User, error) {
	if err := u.db.Where(model.User{Email: user.Email, Password: user.Password}).Take(&user).Error; err != nil {
		return nil, err
	}

	//-----
	// JWT
	//-----

	// Create token
	token, err := util.CreateJwt(user)
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		db,
	}
}
