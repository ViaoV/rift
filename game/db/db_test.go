package db

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestDBConnection(t *testing.T) {
	var err error
	session, err := mgo.Dial(DatabaseHost)
	defer session.Close()
	if err != nil {
		t.Fatalf("Database connection error %s", err.Error())
	}
}

func TestGetSession(t *testing.T) {
	session := GetSession()
	if session == nil {
		t.Fatal("session is nil")
	}
}

func TestGetDatabase(t *testing.T) {
	ses, db := GetDatabase()
	defer ses.Close()
	if db == nil {
		t.Fatal("database is nil")
	}
}

func TestCollection(t *testing.T) {
	ses, c := GetCollection("test")
	defer ses.Close()
	if c == nil {
		t.Fatal("collection is nil")
	}
}
