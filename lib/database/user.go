package database

import (
	"yukevent/config"
	"yukevent/model"
)

func GetAllUser() (*[]model.User, error) {

	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return &[]model.User{}, err
	}
	return &users, nil
}

func GetOneUserByID(id int) (*model.User, error) {

	var user model.User

	err := config.DB.First(&user, id).Error

	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func CreateUser(user model.User) (*model.User, error) {

	err := config.DB.Save(&user).Error

	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func UpdateUser(id int, user model.User) (*model.User, error) {

	err := config.DB.Find(&user, id).Updates(&user).Error
	if err != nil {
		return &model.User{}, err
	}
	return &user, err
}

func DeleteUser(id int, user model.User) (*model.User, error) {

	err := config.DB.Delete(&user, id).Error
	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func DeleteSoft() (*[]model.User, error) {

	var users []model.User

	err := config.DB.Unscoped().Find(&users).Error

	if err != nil {
		return &[]model.User{}, err
	}
	return &users, nil
}
