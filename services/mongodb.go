package services

import (
	"fmt"

	"authentication/config"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// MongoDatabase is an interface to abstract the Database struct from mgo.v2 (*mgo.Database)
type MongoDatabase interface {
	C(name string) *mgo.Collection
}

// MongoCollection is an interface to abstract the Collection struct from mgo.v2 (*mgo.Collection)
type MongoCollection interface {
	Find(query interface{}) *mgo.Query
	EnsureIndex(index mgo.Index) error
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
}

// MongoSession is an interface to abstract the Session struct from mgo.v2 (*mgo.Session)
type MongoSession interface {
	Dial(url string) (*mgo.Session, error)
	Ping() error
	DB(name string) *mgo.Database
}

// DBService is a struct that contains a pointer to the database instance and to each available collection:
// Products
type DBService struct {
	DB    MongoDatabase
	Users MongoCollection
	Roles MongoCollection
}

// InitDBService inits database connection and assigns pointers to Referrer and Referral collections
func (db *DBService) InitDBService() {
	// get a mongo session
	s, err := mgo.Dial(appenv.MustGetEnv(appenv.MONGO_HOST))
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	// assign database
	db.DB = s.DB(appenv.MustGetEnv(appenv.MONGO_DATABASE))

	// assign collections
	db.Users = db.DB.C("users")

	// assign collections
	db.Roles = db.DB.C("roles")

	// set indexes
	usersIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	rolesIndex := mgo.Index{
		Key:        []string{"role, level"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	// users index
	err = db.Users.EnsureIndex(usersIndex)

	// users index
	err = db.Roles.EnsureIndex(rolesIndex)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo database successfully.")
}
