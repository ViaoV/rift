package things

import (
	"os"
	"rift/game/db"
	"testing"
)

func TestMain(m *testing.M) {
	db.DatabaseName = "rift_test"

	result := m.Run()

	db.GetDatabase().DropDatabase()

	os.Exit(result)
}
