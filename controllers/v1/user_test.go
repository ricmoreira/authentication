package controllers

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/util/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// stub UserService behaviour
type MockUserService struct{}

// mocked behaviour for CreateOne
func (us *MockUserService) CreateOne(uReq *mrequest.UserCreate) (*models.User, *mresponse.ErrorResponse) {
	uRes := models.User{}

	// validate request
	err := errors.ValidateRequest(uReq)
	if err != nil {
		return nil, err
	}

	uRes.ID = "some-object-id"
	uRes.Username = uReq.Username
	uRes.Email = uReq.Email
	uRes.Password = uReq.Password

	uRes.Roles = make([]*models.Role, len(uReq.Roles))
	copy(uRes.Roles, uReq.Roles)

	// save user to database
	// (...)

	return &uRes, nil
}

// mocked behaviour for ReadOne
func (ps *MockUserService) ReadOne(p *mrequest.UserRead) (*models.User, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for UpdateOne
func (ps *MockUserService) UpdateOne(p *mrequest.UserUpdate) (*models.User, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for DeleteOne
func (ps *MockUserService) DeleteOne(p *mrequest.UserDelete) (*models.User, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func TestCreateUserAction(t *testing.T) {

	// Mock the server

	// Switch to test mode in order to don't get such noisy output
	gin.SetMode(gin.TestMode)

	ups := MockUserService{}

	uc := UserController{
		UserService: &ups,
	}

	r := gin.Default()

	r.POST("/api/v1/user", uc.CreateAction)

	// TEST SUCCESS

	// Mock a request
	body := mrequest.UserCreate{
		Username: "some-username",
		Email:    "some_email@email.com",
		Password: "some-password",
		Roles: []*models.Role{
			&models.Role{
				Role:  "ADMIN",
				Level: 0,
			},
			&models.Role{
				Role:  "POS",
				Level: 0,
			},
		},
	}

	jsonValue, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	// TEST ERROR ON REQUEST

	// Mock a request
	body = mrequest.UserCreate{}
	jsonValue, _ = json.Marshal(body)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w = httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != 400 {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	// TEST ERROR ON USER SERVICE

	// Mock a request
	body = mrequest.UserCreate{
		Username: "some-username",
		Email:    "error-causing-email",
		Password: "some-password",
		Roles: []*models.Role{
			&models.Role{
				Role:  "ADMIN",
				Level: 0,
			},
			&models.Role{
				Role:  "POS",
				Level: 0,
			},
		},
	}
	jsonValue, _ = json.Marshal(body)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w = httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != 400 {
		t.Fatalf("Expected to get status %d but instead got %d\n", 400, w.Code)
	}
}
