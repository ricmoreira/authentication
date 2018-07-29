package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	Username string `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string `json:"email" valid:"email~Invalid email,required~Field email cannot be empty or is missing" bson:"email"`
	Password string `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []Role `json:"roles" bson:"roles"`
}
