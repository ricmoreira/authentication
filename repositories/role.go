package repositories

import (
	"authentication/models"
	"authentication/models/request"
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// RoleRepository performs CRUD operations on roles resource
type RoleRepository struct {
	col MongoCollection
}

// NewRoleRepository is the constructor for RoleRepository
func NewRoleRepository(db *DBCollections) *RoleRepository {
	return &RoleRepository{col: db.Role}
}

// CreateOne saves provided model instance to database
func (this *RoleRepository) CreateOne(request *mrequest.RoleCreate) (*mongo.InsertOneResult, error) {

	// save role to database
	return this.col.InsertOne(context.Background(), request)
}

// ReadOne returns a role based on role and level sent in request
// TODO: implement better query based on full request and not only the role and the level
func (this *RoleRepository) ReadOne(p *mrequest.RoleRead) (*models.Role, error) {
	result := this.col.FindOne(
		context.Background(),
		bson.NewDocument(bson.EC.String("role", p.Role), bson.EC.Int32("level", int32(p.Level))),
	)

	r := models.Role{}
	err := result.Decode(r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

// TODO: implement
func (this *RoleRepository) UpdateOne(p *mrequest.RoleUpdate) (*models.Role, error) {
	return nil, nil
}

// TODO: implement
func (this *RoleRepository) DeleteOne(p *mrequest.RoleDelete) (*models.Role, error) {
	return nil, nil
}
