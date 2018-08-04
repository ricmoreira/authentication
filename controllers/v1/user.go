package controllers

import (
	"authentication/models/request"
	"authentication/services"
	"authentication/util/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// UserController represents the controller for operating on the users resource
	UserController struct {
		UserService  services.UserServiceContract
		TokenService services.TokenServiceContract
	}
)

// NewUserController is the contructor of UserController
func NewUserController(us *services.UserService, ts *services.TokenService) *UserController {
	return &UserController{
		UserService:  us,
		TokenService: ts,
	}
}

// CreateAction creates a new user
func (this UserController) CreateAction(c *gin.Context) {

	uReq := mrequest.UserCreate{}

	fmt.Println(uReq)
	json.NewDecoder(c.Request.Body).Decode(&uReq)

	e := errors.ValidateRequest(&uReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	uRes, err := this.UserService.CreateOne(&uReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, uRes)
}

// LoginAction sends a JWT token on a success login request
func (this UserController) LoginAction(c *gin.Context) {
	uReq := mrequest.UserLogin{}
	json.NewDecoder(c.Request.Body).Decode(&uReq)

	e := errors.ValidateRequest(&uReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	uRes, cookie, err := this.TokenService.GenerateToken(&uReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, uRes)
}
