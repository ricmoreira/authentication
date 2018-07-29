package server

import (
	"authentication/config"
	"authentication/controllers/v1"
	"authentication/handlers"
	"authentication/middleware"

	"github.com/gin-gonic/gin"
)

// Server is the http layer for role and user resource
type Server struct {
	config         *config.Config
	roleController *controllers.RoleController
	userController *controllers.UserController
	middleware     *middleware.Middleware
	handlers       *handlers.HttpHandlers
}

// NewServer is the Server constructor
func NewServer(cf *config.Config,
	rc *controllers.RoleController,
	uc *controllers.UserController,
	mid *middleware.Middleware,
	hand *handlers.HttpHandlers) *Server {

	return &Server{
		config:         cf,
		roleController: rc,
		userController: uc,
		middleware:     mid,
		handlers:       hand,
	}
}

// Run loads server with its routes and starts the server
func (s *Server) Run() {
	// Instantiate a new router
	r := gin.Default()

	// cors
	r.Use(*s.middleware.Cors())

	// generic routes
	r.HandleMethodNotAllowed = false
	r.NoRoute(s.handlers.NotFound)

	// User resource
	userApi := r.Group("/api/v1/user")
	{
		// Create a new user
		userApi.POST("", s.userController.CreateAction)

		// Emit a jwt token
		userApi.POST("/jwt", s.userController.LoginAction)
	}

	// Roles resource
	roleApi := r.Group("/api/v1/role")
	{
		// Create a new role
		roleApi.POST("", s.roleController.CreateAction)
	}

	// Fire up the server
	r.Run(s.config.Host)
}
