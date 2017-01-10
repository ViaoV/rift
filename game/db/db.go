package db

import (
	logging "github.com/op/go-logging"
	mgo "gopkg.in/mgo.v2"
)

var (
	// DatabaseName is the Mongo Database to read game data from
	DatabaseName = "rift"
	// DatabaseHost is the server address for the database
	DatabaseHost = "localhost"
	log          = logging.MustGetLogger("db")
)

// GetCollection returns a Mongo collection object
func GetCollection(collectionName string) *mgo.Collection {
	session, err := mgo.Dial(DatabaseHost)
	if err != nil {
		log.Criticalf("Database connection error: %s", err.Error())
	}
	db := session.DB(DatabaseName)
	c := db.C(collectionName)
	return c
}

// GetDatabase returns a pointer to the database
func GetDatabase() *mgo.Database {
	db, err := mgo.Dial(DatabaseHost)
	if err != nil {
		log.Criticalf("Database error: %v", err)
		return nil
	}
	return db.DB(DatabaseName)
}

// GetSession returns a mongo session
func GetSession() *mgo.Session {
	session, err := mgo.Dial(DatabaseHost)
	if err != nil {
		log.Criticalf("Database connection error: %s", err.Error())
	}
	return session
}
