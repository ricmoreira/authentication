package controllers

import (
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/util/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// stub RoleService behaviour
type MockRoleService struct{}

// mocked behaviour for CreateOne
func (us *MockRoleService) CreateOne(uReq *mrequest.RoleCreate) (*models.Role, *mresponse.ErrorResponse) {
	rRes := models.Role{}

	// validate request
	err := errors.ValidateRequest(uReq)
	if err != nil {
		return nil, err
	}

	rRes.Role = uReq.Role
	rRes.Level = uReq.Level

	// save role to database
	// (...)

	return &rRes, nil
}

// mocked behaviour for ReadOne
func (ps *MockRoleService) ReadOne(p *mrequest.RoleRead) (*models.Role, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for UpdateOne
func (ps *MockRoleService) UpdateOne(p *mrequest.RoleUpdate) (*models.Role, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for DeleteOne
func (ps *MockRoleService) DeleteOne(p *mrequest.RoleDelete) (*models.Role, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func TestCreateRoleAction(t *testing.T) {

	// Mock the server

	// Switch to test mode in order to don't get such noisy output
	gin.SetMode(gin.TestMode)

	ups := MockRoleService{}

	rc := RoleController{
		RoleService: &ups,
	}

	r := gin.Default()

	r.POST("/api/v1/role", rc.CreateAction)

	// TEST SUCCESS

	// Mock a request
	body := mrequest.RoleCreate{
		Role:  "ADMIN",
		Level: 0,
	}

	jsonValue, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/role", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		bodyString := string(bodyBytes)

		t.Fatalf("Expected to get status %d but instead got %d\nResponse body:\n%s", http.StatusOK, w.Code, bodyString)
	}

	// TEST ERROR ON REQUEST

	// Mock a request
	body = mrequest.RoleCreate{}
	jsonValue, _ = json.Marshal(body)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/role", bytes.NewBuffer(jsonValue))
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

	// TEST ERROR ON ROLE SERVICE

	// Mock a request
	body = mrequest.RoleCreate{
		Level: 0,
	}
	jsonValue, _ = json.Marshal(body)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/role", bytes.NewBuffer(jsonValue))
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
