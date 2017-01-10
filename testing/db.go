package testhelper

import (
	"os"
	"rift/game/config"

	mgo "gopkg.in/mgo.v2"
)

type (
	WithSessionHandler func(session mgo.Session)
)

// SetupDB helpers manage testing and tearing down the database
func SetupDB(handler WithSessionHandler) {
	config.DatabaseName = "rift_test"
	session, err := mgo.Dial(config.DatabaseHost)
	defer session.Close()
	handler(session)
	session.DB("rift_test").DropDatabase()
	os.Exit(result)
}
