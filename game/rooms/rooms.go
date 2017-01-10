package rooms

import (
	"rift/game/config"
	"rift/game/lang"

	"github.com/emirpasic/gods/sets/hashset"

	logging "github.com/op/go-logging"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// RoomOccupants is an interface for objects that can occupy rooms
	RoomOccupants interface {
		ShortNames() []string
	}
	// RoomProperties is a collection of string to string values that can
	// be attached to a room
	RoomProperties map[string]string
	// Room is the basic room configuration
	Room struct {
		ID          bson.ObjectId `bson:"_id"`
		Title       string
		Description string
		RoomExits   []RoomExit
		Flags       *hashset.Set
		Properties  RoomProperties
		Occupants   RoomOccupants
	}
)

var log = logging.MustGetLogger("rooms")

// LoadRoom loads a room from the data files
func LoadRoom(id bson.ObjectId) (*Room, error) {
	session, _ := mgo.Dial(config.DatabaseHost)
	defer session.Close()
	c := session.DB(config.DatabaseName).C("rooms")
	var room Room
	err := c.Find(bson.M{"_id": id}).One(&room)
	return &room, err
}

// AlsoHere returns a formatted sentence of characters in the room.
func (room *Room) AlsoHere() string {
	return lang.MakeSentence(room.Occupants.ShortNames())
}

// NewRoom constructs a new Room
func NewRoom() *Room {
	return &Room{
		ID:    bson.NewObjectId(),
		Flags: hashset.New(),
	}
}
