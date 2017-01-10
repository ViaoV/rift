package things

import (
	"testing"
)

func TestShortName(t *testing.T) {
	form := NewItemForm()
	form.Noun = "sword"
	form.PrimaryAdjective = "short"
	if form.ShortName() != "short sword" {
		t.Fatalf("Short name should be 'short sword', was %s", form.ShortName())
	}
}

// func TestNewForm(t *testing.T) {
// 	form := NewItemForm()
// 	if form.Flags == nil {
// 		t.Fatal("Flags was not initialized")
// 	}
// }

func TestIsWearable(t *testing.T) {
	form := NewItemForm()
	if form.IsWearable() != false {
		t.Error("IsWearable should return false")
	}
	form.WearableLocation = InventorySlotHead
	if form.IsWearable() != true {
		t.Error("IsWearable should return true")
	}
	form.WearableLocation = InventorySlotNone
	if form.IsWearable() != false {
		t.Error("IsWearable should return false")
	}
}

func TestGetWearableLocation(t *testing.T) {
	form := NewItemForm()
	form.WearableLocation = InventorySlotHead
	if form.GetWearableLocation() != InventorySlotHead {
		t.Errorf("GetWearableLocation() should return %s, was %s",
			InventorySlotHead, form.GetWearableLocation())
	}
}
