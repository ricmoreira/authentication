package repositories

import (
	"authentication/helper"
	"authentication/models"
	"authentication/models/request"
	"context"
	"errors"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// UserRepository performs CRUD operations on users resource
type UserRepository struct {
	users MongoCollection
	roles MongoCollection
}

// NewUserRepository is the constructor for UserRepository
func NewUserRepository(db *DBCollections) *UserRepository {
	return &UserRepository{users: db.User, roles: db.Role}
}

// CreateOne saves provided model instance to database
func (this *UserRepository) CreateOne(request *mrequest.UserCreate) (*mongo.InsertOneResult, error) {

	// encript password
	password, e := helper.Encrypt(request.Password)
	if e != nil {
		return nil, e
	}

	request.Password = password

	// validate roles exist in db
	err := this.ValidateRoles(request.Roles)
	if err != nil {
		return nil, err
	}

	// save user to database
	return this.users.InsertOne(context.Background(), request)
}

// ReadOne returns a user based on username sent in request
// TODO: implement better query based on full request and not only the username
func (this *UserRepository) ReadOne(req *mrequest.UserRead) (*models.User, error) {
	result := this.users.FindOne(
		context.Background(),
		bson.NewDocument(bson.EC.String("username", req.Username)),
	)


	u := models.User{}
	err := result.Decode(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// TODO: implement
func (this *UserRepository) UpdateOne(p *mrequest.UserUpdate) (*models.User, error) {
	return nil, nil
}

// TODO: implement
func (this *UserRepository) DeleteOne(p *mrequest.UserDelete) (*models.User, error) {
	return nil, nil
}

// ValidateRoles check if provided roles exist in roles collection
func (this *UserRepository) ValidateRoles(roles []*models.Role) error {

	queue := make(chan RoleResult, len(roles))

	for _, role := range roles {
		rr := RoleResult{
			Role:  role.Role,
			Level: role.Level,
		}

		rr.Result = this.roles.FindOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.String("role", role.Role),
				bson.EC.Int32("level", int32(role.Level)),
			))

		queue <- rr
	}

	close(queue)

	for elem := range queue {
	
		var result interface{}
		err := elem.Result.Decode(result)

		if err != nil {
			return errors.New("Role: " + elem.Role + " with Level: " + strconv.Itoa(int(elem.Level)) + " does not exist. ")
		}
	}

	return nil
}

// RoleResult is a struct helper for mapping a *mongo.DocumentResult with provided findOne() request 
type RoleResult struct {
	Role   string
	Level  int32
	Result *mongo.DocumentResult
}
