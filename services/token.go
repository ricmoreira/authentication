package services

import (
	"authentication/config"
	"authentication/helper"
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/repositories"
	"authentication/util/errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenServiceContract is the abstraction for service layer on token resource
type TokenServiceContract interface {
	GenerateToken(request *mrequest.UserLogin) (*models.Token, *http.Cookie, *mresponse.ErrorResponse)
}

// TokenService provides custom JWT tokens based on request resource
type TokenService struct {
	userRepository *repositories.UserRepository
	config         *config.Config
}

// NewTokenService is the constructor of TokenService
func NewTokenService(rr *repositories.UserRepository, config *config.Config) *TokenService {
	return &TokenService{
		userRepository: rr,
		config: config,
	}
}

// CreateOne saves provided model instance to database
func (this *TokenService) GenerateToken(request *mrequest.UserLogin) (*models.Token, *http.Cookie, *mresponse.ErrorResponse) {
	// get from env secret key and expiration time
	secretKey := os.Getenv("JWT_SECRET_KEY")
	envExpir := os.Getenv("JWT_TOKEN_MIN_EXPIRE")
	intExpir, err := strconv.ParseInt(envExpir, 10, 64)
	if err != nil {
		errR := errors.HandleErrorResponse(errors.UNKNOWN_ERROR, nil, err.Error())
		return nil, nil, errR
	}

	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, nil, e
	}

	// find user in database
	userRequest := mrequest.UserRead{
		Username: request.Username,
	}

	u, err := this.userRepository.ReadOne(&userRequest)
	if err != nil {
		e = errors.HandleErrorResponse(errors.UNAUTHORIZED, nil, err.Error())
		return nil, nil, e
	}

	// match password
	err = helper.Compare(u.Password, request.Password)
	if err != nil {
		e = errors.HandleErrorResponse(errors.UNAUTHORIZED, nil, err.Error())
		return nil, nil, e
	}

	//build claims

	claims := models.Claims{}

	// assign custom claims
	claims.Username = u.Username
	claims.Roles = make([]*models.Role, len(u.Roles))
	copy(claims.Roles, u.Roles)

	// prepare token expiration
	now := time.Now()
	expir := now.Add(time.Minute * time.Duration(intExpir))

	// assign values to standard claims
	claims.Issuer = "authentication"
	claims.Audience = "ecommerce_admin"
	claims.ExpiresAt = expir.Unix()
	claims.IssuedAt = now.Unix()

	// create token & sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))

	if err != nil {
		errR := errors.HandleErrorResponse(errors.UNKNOWN_ERROR, nil, err.Error())
		return nil, nil, errR
	}

	res := models.Token{
		Token: t,
	}

	// set token to cookie
	encodedToken := url.QueryEscape(t)

	cookie := http.Cookie{
		Name:     "JWT",
		Value:    encodedToken,
		Expires:  expir,
		Path:     "/",
		Domain:   this.config.CookieDomain,
		Secure:   true,
		HttpOnly: true,
	}

	return &res, &cookie, nil
}
