package controllers

import (
	"authentication/models/request"
	"authentication/services"
	"authentication/util/errors"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type (
	// RoleController represents the controller for operating on the roles resource
	RoleController struct {
		RoleService services.RoleService
	}
)

// CreateAction creates a new role
func (rc RoleController) CreateAction(c *gin.Context) {
	rReq := mrequest.RoleCreate{}
	json.NewDecoder(c.Request.Body).Decode(&rReq)

	e := errors.ValidateRequest(&rReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	rRes, err := rc.RoleService.CreateOne(&rReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, rRes)
}
