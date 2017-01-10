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
	// Session is the currently active database connection
	session *mgo.Session
)

// GetCollection returns a Mongo collection object
func GetCollection(collectionName string) *mgo.Collection {
	session = GetSession()
	db := session.DB(DatabaseName)
	c := db.C(collectionName)
	return c
}

// GetDatabase returns a pointer to the database
func GetDatabase() *mgo.Database {
	return GetSession().DB(DatabaseName)
}

// GetSession returns a mongo session
func GetSession() *mgo.Session {
	var err error
	if session != nil {
		return session
	}
	session, err = mgo.Dial(DatabaseHost)
	if err != nil {
		log.Criticalf("Database connection error: %s", err.Error())
	}
	return session
}
