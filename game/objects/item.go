package things

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

type (
	// Item represnets an interactable/equipable thing
	Item struct {
		Noun             string
		PrimaryAdjective string
		FullName         string
		Look             string
		Read             string
		WearableLocation string
		Flags            *hashset.Set
	}
)

// ShortName is the displable name of the item
func (item *Item) ShortName() string {
	return fmt.Sprintf("%s %s", item.PrimaryAdjective, item.Noun)
}

// IsWearable returns true if the item is equipable
func (item *Item) IsWearable() bool {
	if item.WearableLocation == "" {
		return false
	}
	return true
}

// GetWearableLocation returns the wearbale location of the item
func (item *Item) GetWearableLocation() string {
	return item.WearableLocation
}

// NewItem initializes a new item
func NewItem() Item {
	return Item{
		Flags: hashset.New(),
	}
}
