package model

import "github.com/google/uuid"

type RegisterUser struct {
	ID       uuid.UUID
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	UserName string    `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}