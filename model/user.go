package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
	Sex      string `json:"sex" form:"sex"`
	ClientID int    `json:"client_id" form:"client_id"`
}
