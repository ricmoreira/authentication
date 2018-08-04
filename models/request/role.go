package mrequest

type RoleCreate struct {
	Role  string `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level int32    `json:"level" bson:"level"`
}

type RoleRead struct {
	Role  string `json:"role" bson:"role"`
	Level int32    `json:"level" bson:"level"`
}

type RoleUpdate struct {
	Role  string `json:"role" bson:"role"`
	Level int32    `json:"level" bson:"level"`
}

type RoleDelete struct {
	Role  string `json:"role" valid:"required~Field role cannot be empty" bson:"role"`
	Level int32    `json:"level" valid:"required~Field level cannot be empty" bson:"level"`
}
