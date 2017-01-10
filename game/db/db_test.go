package db

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestDBConnection(t *testing.T) {
	var err error
	session, err = mgo.Dial(DatabaseHost)
	if err != nil {
		t.Fatalf("Database connection error %s", err.Error())
	}
}

func TestGetSession(t *testing.T) {
	session = GetSession()
	if session == nil {
		t.Fatal("session is nil")
	}
}

func TestGetDatabase(t *testing.T) {
	db := GetDatabase()
	if db == nil {
		t.Fatal("database is nil")
	}
}

func TestCollection(t *testing.T) {
	c := GetCollection("test")
	if c == nil {
		t.Fatal("collection is nil")
	}
}

func TestFailureLogging(t *testing.T) {
	DatabaseHost = "169.1.1.1"
	session = nil
	ses := GetSession()
	if ses != nil {
		t.Fatal("Connection failed in an unexpected way")
	}
}
