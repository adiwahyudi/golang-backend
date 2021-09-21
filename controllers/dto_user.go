package controllers

import (
	"time"
	"yukevent/model"
)

type RequestUser struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Sex      string `json:"sex" form:"sex"`
	ClientID int    `json:"client_id" form:"client_id"`
	Email    string `json:"email" form:"email"`
}

func (req *RequestUser) toModel() model.User {
	return model.User{
		ID:       req.ID,
		Name:     req.Name,
		Age:      req.Age,
		Sex:      req.Sex,
		ClientID: req.ClientID,
		Email:    req.Email,
	}
}

type ResponseUser struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Age       int       `json:"age" form:"age"`
	Sex       string    `json:"sex" form:"sex"`
	ClientID  int       `json:"client_id" form:"client_id"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

func newResponse(mdlUsers model.User) ResponseUser {
	return ResponseUser{
		ID:        mdlUsers.ID,
		Name:      mdlUsers.Name,
		Age:       mdlUsers.Age,
		Sex:       mdlUsers.Sex,
		ClientID:  mdlUsers.ClientID,
		Email:     mdlUsers.Email,
		CreatedAt: mdlUsers.CreatedAt,
		UpdatedAt: mdlUsers.UpdatedAt,
		DeletedAt: mdlUsers.DeletedAt.Time,
	}
}

func newResponseArray(mdlUsers []model.User) []ResponseUser {
	var resp []ResponseUser

	for _, value := range mdlUsers {
		resp = append(resp, newResponse(value))
	}

	return resp
}
