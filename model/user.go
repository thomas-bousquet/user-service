package model

import (
	"time"
)

type User struct {
	Id string `json:"id" bson:"_id"`
	Firstname string `json:"first_name" bson:"first_name" validate:"required"`
	Lastname string `json:"last_name" bson:"last_name" validate:"required"`
	Email string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}

func NewUser(firstname string, lastname string, email string, password string) User {
	return User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password: password,
		CreatedAt: time.Now(),
	}
}
