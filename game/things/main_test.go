package things

import (
	"os"
	"rift/game/db"
	"testing"
)

func TestMain(m *testing.M) {
	db.DatabaseName = "rift_test"
	session := db.GetSession()

	result := m.Run()

	session.DB("rift_test").DropDatabase()
	os.Exit(result)
}
