package models

type Role struct {
	ID    string `json:"id,omitempty" bson:"_id"`
	Role  string `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level int    `json:"level" valid:"required~Field level cannot be empty or is missing" bson:"level"`
}
