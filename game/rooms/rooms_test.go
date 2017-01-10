package rooms

import (
	"os"
	"rift/game/config"
	"rift/game/entities"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	northRoom *Room
	southRoom *Room
)

func TestMain(m *testing.M) {
	config.DatabaseName = "rift_test"
	session, err := mgo.Dial(config.DatabaseHost)
	if err != nil {
		panic(err)
	}
	c := session.DB("rift_test").C("rooms")
	defer session.Close()

	northRoom = NewRoom()
	northRoom.Title = "The room to the north"
	northRoom.Description = "Room Description"
	northRoom.Properties = RoomProperties{"test_prop": "test_value"}
	northRoom.Flags.Add("test")
	northRoom.RoomExits = append(northRoom.RoomExits, RoomExit{
		Description:   "south",
		DestinationID: bson.NewObjectId(),
	})

	southRoom = NewRoom()
	northRoom.Title = "The room to the south"
	northRoom.Description = "Room Description"
	northRoom.Properties = RoomProperties{"test_prop": "test_value"}
	northRoom.Flags.Add("test")
	northRoom.RoomExits = append(northRoom.RoomExits, RoomExit{
		Description:   "north",
		DestinationID: northRoom.ID,
	})
	northRoom.RoomExits[0].DestinationID = southRoom.ID

	err = c.Insert(northRoom, southRoom)

	if err != nil {
		panic(err)
	}

	result := m.Run()
	if err != nil {
		os.Exit(-1)
	}
	session.DB("rift_test").DropDatabase()
	os.Exit(result)
}

func TestLoadRoom(t *testing.T) {
	room, err := LoadRoom(northRoom.ID)
	if err != nil {
		t.Fatal(err)
	}
	if room == nil {
		t.Fatal("Room did not load")
	}
	if room.Title != northRoom.Title {
		t.Fatal("Room title not found")
	}
	if room.Properties["test_prop"] != northRoom.Properties["test_prop"] {
		t.Fatal("Room properties did not load")
	}
	if room.Flags.Contains("test") {
		t.Fatal("Room flags did not load")
	}
}

func TestAlsoHere(t *testing.T) {
	room := Room{
		ID: bson.NewObjectId(),
		Occupants: entities.CharacterCollection{
			&entities.Character{Name: "Joe"},
			&entities.Character{Name: "Jim"},
		},
	}
	if room.AlsoHere() != "Joe and Jim" {
		t.Errorf("Also here badly formed: %s", room.AlsoHere())
	}
}
