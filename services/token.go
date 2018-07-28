package services

import (
	"authentication/config"
	"authentication/models"
	"authentication/models/request"
	"authentication/models/response"
	"authentication/util/errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenService provides custom JWT tokens based on request resource
type MongoTokenService struct {
	DBService *DBService
}

// CreateOne saves provided model instance to database
func (ts *MongoTokenService) GenerateToken(request *mrequest.UserLogin) (*models.Token, *http.Cookie, *mresponse.ErrorResponse) {
	// get from env secret key and expiration time
	secretKey := os.Getenv("JWT_SECRET_KEY")
	envExpir := os.Getenv("JWT_TOKEN_MIN_EXPIRE")
	intExpir, err := strconv.ParseInt(envExpir, 10, 64)
	if err != nil {
		errR := errors.HandleErrorResponse(errors.UNKNOWN_ERROR, nil, err.Error())
		return nil, nil, errR
	}

	u := models.User{}

	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, nil, e
	}

	// find user in database
	err = ts.DBService.Users.Find(bson.M{"username": request.Username}).One(&u)
	if err != nil {
		e = errors.HandleErrorResponse(errors.UNAUTHORIZED, nil, err.Error())
		return nil, nil, e
	}

	// match password
	err = Compare(u.Password, request.Password)
	if err != nil {
		e = errors.HandleErrorResponse(errors.UNAUTHORIZED, nil, err.Error())
		return nil, nil, e
	}

	//build claims

	claims := models.Claims{}

	// assign custom claims
	claims.Username = u.Username
	claims.Roles = make([]models.Role, len(u.Roles))
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
		Domain:   appenv.MustGetEnv(appenv.COOKIE_DOMAIN),
		Secure:   true,
		HttpOnly: true,
	}

	return &res, &cookie, nil
}
