package rooms

import "gopkg.in/mgo.v2/bson"

type (
	// RoomExit is a link from one room to another room
	RoomExit struct {
		Description   string
		DestinationID bson.ObjectId
	}
)

//Destination returns the loaded Room
func (exit *RoomExit) Destination() *Room {
	room, _ := LoadRoom(exit.DestinationID)
	return room
}
