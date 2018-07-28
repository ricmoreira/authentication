package controllers

import (
	"authentication/models/request"
	"authentication/services"
	"authentication/util/errors"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// UserController represents the controller for operating on the users resource
	UserController struct {
		UserService  services.UserService
		TokenService *services.MongoTokenService
	}
)

// CreateAction creates a new user
func (uc UserController) CreateAction(c *gin.Context) {
	uReq := mrequest.UserCreate{}
	json.NewDecoder(c.Request.Body).Decode(&uReq)

	e := errors.ValidateRequest(&uReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	uRes, err := uc.UserService.CreateOne(&uReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, uRes)
}

// LoginAction sends a JWT token on a success login request
func (uc UserController) LoginAction(c *gin.Context) {
	uReq := mrequest.UserLogin{}
	json.NewDecoder(c.Request.Body).Decode(&uReq)

	e := errors.ValidateRequest(&uReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	uRes, cookie, err := uc.TokenService.GenerateToken(&uReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, uRes)
}
