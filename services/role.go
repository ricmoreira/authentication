package services

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/util/errors"

	"gopkg.in/mgo.v2/bson"
)

// RoleService performs CRUD operations on roles resource
type RoleService interface {
	CreateOne(p *mrequest.RoleCreate) (*models.Role, *mresponse.ErrorResponse)
	ReadOne(p *mrequest.RoleRead) (*models.Role, *mresponse.ErrorResponse)
	UpdateOne(p *mrequest.RoleUpdate) (*models.Role, *mresponse.ErrorResponse)
	DeleteOne(p *mrequest.RoleDelete) (*models.Role, *mresponse.ErrorResponse)
}

// MongoReferrerService is a struct that contains a DBService pointer that exposes the referral collection and its database
type MongoRoleService struct {
	DBService *DBService
}

// CreateOne saves provided model instance to database
func (mus *MongoRoleService) CreateOne(request *mrequest.RoleCreate) (*models.Role, *mresponse.ErrorResponse) {

	r := models.Role{}

	// validate request
	err := errors.ValidateRequest(request)
	if err != nil {
		return nil, err
	}

	r.ID = bson.NewObjectId()
	r.Role = request.Role
	r.Level = request.Level

	// save role to database
	if err := mus.DBService.Roles.Insert(r); err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	return &r, nil
}

// TODO: implement
func (mus *MongoRoleService) ReadOne(p *mrequest.RoleRead) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (mus *MongoRoleService) UpdateOne(p *mrequest.RoleUpdate) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (mus *MongoRoleService) DeleteOne(p *mrequest.RoleDelete) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}
