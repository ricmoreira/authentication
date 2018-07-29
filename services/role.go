package services

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/repositories"
	"authentication/util/errors"
)

// RoleServiceContract is the abstraction for service layer on role resource
type RoleServiceContract interface {
	CreateOne(p *mrequest.RoleCreate) (*models.Role, *mresponse.ErrorResponse)
	ReadOne(p *mrequest.RoleRead) (*models.Role, *mresponse.ErrorResponse)
	UpdateOne(p *mrequest.RoleUpdate) (*models.Role, *mresponse.ErrorResponse)
	DeleteOne(p *mrequest.RoleDelete) (*models.Role, *mresponse.ErrorResponse)
}

// RoleService is the layer between http client and repository for role resource
type RoleService struct {
	repository *repositories.RoleRepository
}

// NewRoleService is the constructor of RoleService
func NewRoleService(rr *repositories.RoleRepository) *RoleService {
	return &RoleService{
		repository: rr,
	}
}

// CreateOne saves provided model instance to database
func (this *RoleService) CreateOne(request *mrequest.RoleCreate) (*models.Role, *mresponse.ErrorResponse) {

	// validate request
	err := errors.ValidateRequest(request)
	if err != nil {
		return nil, err
	}

	r, e := this.repository.CreateOne(request)

	if e != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, e.Error())
		return nil, errR
	}

	return r, nil
}

// TODO: implement
func (this *RoleService) ReadOne(p *mrequest.RoleRead) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *RoleService) UpdateOne(p *mrequest.RoleUpdate) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (me *RoleService) DeleteOne(p *mrequest.RoleDelete) (*models.Role, *mresponse.ErrorResponse) {
	return nil, nil
}
