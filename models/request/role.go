package mrequest

type RoleCreate struct {
	Role  string `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level int    `json:"level" bson:"level"`
}

type RoleRead struct {
	ID    string `json:"id,omitempty" bson:"_id"`
	Role  string `json:"role" bson:"role"`
	Level int    `json:"level" bson:"level"`
}

type RoleUpdate struct {
	ID    string `json:"id,omitempty" valid:"required~Field ID cannot be empty" bson:"_id"`
	Role  string `json:"role" bson:"role"`
	Level int    `json:"level" bson:"level"`
}

type RoleDelete struct {
	ID string `json:"id,omitempty" bson:"_id"`
}
