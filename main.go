package main

import (
	"authentication/config"
	"authentication/controllers/v1"
	"authentication/services"
	"authentication/util/errors"
	"authentication/models/response"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// check all mandatory environment variables
	appenv.CheckMandatoryEnv()

	// Init Services
	// Database Service
	db := services.DBService{}
	db.InitDBService()

	// Users Service
	us := services.MongoUserService{
		DBService: &db,
	}

	// Token Service
	ts := services.MongoTokenService{
		DBService: &db,
	}

	// Roles Service
	rs := services.MongoRoleService{
		DBService: &db,
	}

	// Get a UserController instance
	uc := controllers.UserController{
		UserService:  &us,
		TokenService: &ts,
	}

	// Get a RoleController instance
	rc := controllers.RoleController{
		RoleService:  &rs,
	}

	// Instantiate a new router
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Accept-Encoding", "Accept-Language", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Cache-Control", "Connection",
			"Host", "Origin", "Pragma", "User-Agent", "X-Custom-Header", "access-control-allow-origin", "authorization", "Origin", "Content-Type", "Accept", "Key", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           48 * time.Hour,
	}))
	r.HandleMethodNotAllowed = false
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, errors.HandleErrorResponse(errors.NOT_FOUND, []mresponse.ErrorDetail{}, ""))
	})

	// User resource
	userApi := r.Group("/api/v1/user")
	{
		// Create a new user
		userApi.POST("", uc.CreateAction)

		// Emit a jwt token
		userApi.POST("/jwt", uc.LoginAction)
	}

	// Roles resource
	roleApi := r.Group("/api/v1/role")
	{
		// Create a new role
		roleApi.POST("", rc.CreateAction)
	}

	// Fire up the server
	r.Run(appenv.MustGetEnv(appenv.HOST))
}
