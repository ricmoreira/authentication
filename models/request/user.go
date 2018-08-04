package mrequest

import (
	"authentication/models"
)

var ROLES = map[string]string{
	"ADMIN": "ADMIN",
	"USER":  "USER",
	"POS":  "POS",
}

type UserLogin struct {
	Username string   `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Password string   `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
}

type UserCreate struct {
	Username string   `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string   `json:"email" valid:"email~Invalid email" bson:"email"`
	Password string   `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []*models.Role `json:"roles" bson:"roles"`
}

type UserRead struct {
	Username string   `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string   `json:"email" valid:"email~Invalid email,required~Field email cannot be empty or is missing" bson:"email"`
	Password string   `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []*models.Role `json:"roles" bson:"roles"`
}

type UserUpdate struct {
	Username string   `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string   `json:"email" valid:"email~Invalid email,required~Field email cannot be empty or is missing" bson:"email"`
	Password string   `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []*models.Role `json:"roles" bson:"roles"`
}

type UserDelete struct {
	Username string   `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string   `json:"email" valid:"email~Invalid email,required~Field email cannot be empty or is missing" bson:"email"`
	Password string   `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []*models.Role `json:"roles" bson:"roles"`
}
