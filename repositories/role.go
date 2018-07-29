package repositories

import (
	"authentication/models"
	"authentication/models/request"
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
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
func (this *RoleRepository) CreateOne(request *mrequest.RoleCreate) (*models.Role, error) {

	// save role to database
	result, e := this.col.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("role", request.Role),
			bson.EC.Int32("level", int32(request.Level)),
		))

	if e != nil {
		return nil, e
	}

	r := models.Role{}
	if str, ok := result.InsertedID.(string); ok {
		r.ID = str
	}

	r.Level = request.Level
	r.Role = request.Role

	return &r, nil
}

// TODO: implement
func (this *RoleRepository) ReadOne(p *mrequest.RoleRead) (*models.Role, error) {
	return nil, nil
}

// TODO: implement
func (this *RoleRepository) UpdateOne(p *mrequest.RoleUpdate) (*models.Role, error) {
	return nil, nil
}

// TODO: implement
func (this *RoleRepository) DeleteOne(p *mrequest.RoleDelete) (*models.Role, error) {
	return nil, nil
}
