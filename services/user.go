package services

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/repositories"
	"authentication/util/errors"
)

// UserServiceContract is the abstraction for service layer on roles resource
type UserServiceContract interface {
	CreateOne(p *mrequest.UserCreate) (*models.User, *mresponse.ErrorResponse)
	ReadOne(p *mrequest.UserRead) (*models.User, *mresponse.ErrorResponse)
	UpdateOne(p *mrequest.UserUpdate) (*models.User, *mresponse.ErrorResponse)
	DeleteOne(p *mrequest.UserDelete) (*models.User, *mresponse.ErrorResponse)
}

// UserService is the layer between http client and repository for user resource
type UserService struct {
	repository *repositories.UserRepository
}

// NewUserService is the constructor of UserService
func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{
		repository: ur,
	}
}

// CreateOne saves provided model instance to database
func (this *UserService) CreateOne(request *mrequest.UserCreate) (*models.User, *mresponse.ErrorResponse) {

	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, e
	}

	u, err := this.repository.CreateOne(request)

	if err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	// TODO: implement roles
	// u.Roles = make([]models.Role, len(request.Roles))
	// copy(u.Roles, request.Roles)

	return u, nil
}

// TODO: implement
func (this *UserService) ReadOne(p *mrequest.UserRead) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *UserService) UpdateOne(p *mrequest.UserUpdate) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *UserService) DeleteOne(p *mrequest.UserDelete) (*models.User, *mresponse.ErrorResponse) {
	return nil, nil
}
