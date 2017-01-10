package entities

type (
	// Entity is a base interface for all other entities
	Entity interface {
		RoomDescription() string
	}
)
