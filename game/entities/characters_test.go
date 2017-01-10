package entities

import (
	"os"
	"rift/game/db"
	"testing"
)

var (
	char1 *Character
	char2 *Character
)

func TestMain(m *testing.M) {
	db.DatabaseName = "rift_test"

	ses, c := db.GetCollection("characters")
	defer ses.Close()

	char1 = NewCharacter()
	char1.Name = "Jim"

	char2 = NewCharacter()
	char1.Name = "John"

	err := c.Insert(char1, char2)

	if err != nil {
		panic(err)
	}

	result := m.Run()

	c.DropCollection()
	os.Exit(result)
}

func TestLoadCharacter(t *testing.T) {
	ch, err := LoadCharacter(char1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if ch.Name != char1.Name {
		t.Fatal("Character did not load properly")
	}
}

func TestShortNames(t *testing.T) {
	coll := CharacterCollection{char1, char2}
	shortNames := coll.ShortNames()
	if len(shortNames) != 2 {
		t.Fatalf("ShortNames returned %d, should be %d", len(shortNames), 2)
	}
}

func TestNames(t *testing.T) {
	coll := CharacterCollection{char1, char2}
	names := coll.Names()
	if len(names) != 2 {
		t.Fatalf("ShortNames returned %d, should be %d", len(names), 2)
	}
}
