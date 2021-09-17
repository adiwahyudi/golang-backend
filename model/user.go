package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
}
