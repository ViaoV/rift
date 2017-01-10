package entities

import (
	"testing"
)

type (
	mockInventoryObject struct {
		shortName           string
		isWearable          bool
		getWearableLocation string
	}
)

func (i *mockInventoryObject) ShortName() string {
	return i.shortName
}

func (i *mockInventoryObject) IsWearable() bool {
	return i.isWearable
}

func (i *mockInventoryObject) GetWearableLocation() string {
	return i.getWearableLocation
}

func TestInventory(t *testing.T) {
	item := &mockInventoryObject{}
	char1.Inventory.RightHand = item
}
