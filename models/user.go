package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Username string        `json:"username" valid:"required~Field username cannot be empty" bson:"username"`
	Email    string        `json:"email" valid:"email~Invalid email,required~Field email cannot be empty or is missing" bson:"email"`
	Password string        `json:"password" valid:"runelength(3|50)~Password must have at least 3 characters" bson:"password"`
	Roles    []Role        `json:"roles" bson:"roles"`
}
