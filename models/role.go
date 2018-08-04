package models

type Role struct {
	Role  string `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level int32    `json:"level" valid:"required~Field level cannot be empty or is missing" bson:"level"`
}
