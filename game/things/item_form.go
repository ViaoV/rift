package things

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	// InventorySlotNone is a constant value for inventory slots
	InventorySlotNone = ""
	// InventorySlotHead is a constant value for inventory slots
	InventorySlotHead = "HEAD"
	// InventorySlotRightHand is a constant value for inventory slots
	InventorySlotRightHand = "RIGHTHAND"
	// InventorySlotLeftHand is a constant value for inventory slots
	InventorySlotLeftHand = "LEFTHAND"
	// InventorySlotRing1 is a constant value for inventory slots
	InventorySlotRing1 = "RING1"
	// InventorySlotRing2 is a constant value for inventory slots
	InventorySlotRing2 = "RING2"
)

type (
	// ItemForm represnets an item ItemForm
	ItemForm struct {
		ID               bson.ObjectId
		Noun             string
		PrimaryAdjective string
		FullName         string
		Look             string
		Read             string
		WearableLocation string
		Weapon           bool
		Static           bool
		Container        bool
		ContainerSize    int
		ObjectSize       int
		Weight           int
		Durability       int
		Value            int
	}
)

// ShortName is the displable name of the item
func (item *ItemForm) ShortName() string {
	return fmt.Sprintf("%s %s", item.PrimaryAdjective, item.Noun)
}

// IsWearable returns true if the item is equipable
func (item *ItemForm) IsWearable() bool {
	if item.WearableLocation == "" {
		return false
	}
	return true
}

// GetWearableLocation returns the wearbale location of the item
func (item *ItemForm) GetWearableLocation() string {
	return item.WearableLocation
}

// NewItemForm initializes a new item
func NewItemForm() *ItemForm {
	return &ItemForm{
		ID: bson.NewObjectId(),
	}
}
