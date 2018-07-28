package services

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/util/errors"

	"gopkg.in/mgo.v2/bson"
)

// UserService performs CRUD operations on users resource
type UserService interface {
	CreateOne(p *mrequest.UserCreate) (*models.User, *mresponse.ErrorResponse)
	ReadOne(p *mrequest.UserRead) (*models.User, *mresponse.ErrorResponse)
	UpdateOne(p *mrequest.UserUpdate) (*models.User, *mresponse.ErrorResponse)
	DeleteOne(p *mrequest.UserDelete) (*models.User, *mresponse.ErrorResponse)
}

// MongoReferrerService is a struct that contains a DBService pointer that exposes the referral collection and its database
type MongoUserService struct {
	DBService *DBService
}

// CreateOne saves provided model instance to database
func (mus *MongoUserService) CreateOne(request *mrequest.UserCreate) (*models.User, *mresponse.ErrorResponse) {

	u := models.User{}

	// validate request
	err := errors.ValidateRequest(request)
	if err != nil {
		return nil, err
	}

	u.ID = bson.NewObjectId()
	u.Username = request.Username
	u.Email = request.Email

	// encript and save password
	password, e := Encrypt(request.Password)

	if e != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, e.Error())
		return nil, errR
	}

	u.Password = password

	u.Roles = make([]models.Role, len(request.Roles))
	copy(u.Roles, request.Roles)

	// save user to database
	if err := mus.DBService.Users.Insert(u); err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	return &u, nil
}

// TODO: implement
func (mus *MongoUserService) ReadOne(p *mrequest.UserRead) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (mus *MongoUserService) UpdateOne(p *mrequest.UserUpdate) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (mus *MongoUserService) DeleteOne(p *mrequest.UserDelete) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}
