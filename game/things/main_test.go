package things

import (
	"os"
	"rift/game/db"
	"testing"
)

func TestMain(m *testing.M) {
	db.DatabaseName = "rift_test"
	result := m.Run()

	ses, c := db.GetCollection("items")
	defer ses.Close()
	c.DropCollection()

	os.Exit(result)
}
