package repositories

import (
	"authentication/helper"
	"authentication/models"
	"authentication/models/request"
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
)

// UserRepository performs CRUD operations on users resource
type UserRepository struct {
	col MongoCollection
}

// NewUserRepository is the constructor for UserRepository
func NewUserRepository(db *DBCollections) *UserRepository {
	return &UserRepository{col: db.User}
}

// CreateOne saves provided model instance to database
func (this *UserRepository) CreateOne(request *mrequest.UserCreate) (*models.User, error) {

	// encript password
	password, e := helper.Encrypt(request.Password)
	if e != nil {
		return nil, e
	}

	// save user to database
	result, e := this.col.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("username", request.Username),
			bson.EC.String("email", request.Email),
			bson.EC.String("password", password),
		))

	if e != nil {
		return nil, e
	}

	r := models.User{}
	if str, ok := result.InsertedID.(string); ok {
		r.ID = str
	}

	r.Username = request.Username
	r.Email = request.Email

	return &r, nil
}

// ReadOne returns a user based on username sent in request
// TODO: implement better query based on full request and not only the username
func (this *UserRepository) ReadOne(p *mrequest.UserRead) (*models.User, error) {
	result := this.col.FindOne(
		context.Background(),
		bson.NewDocument(bson.EC.String("username", p.Username)),
	)

	u := models.User{}
	err := result.Decode(u)

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
