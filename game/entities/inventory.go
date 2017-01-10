package entities

type (

	// InventoryObject is an interface that represents equipable objects
	InventoryObject interface {
		ShortName() string
		IsWearable() bool
		GetWearableLocation() string
	}

	// Inventory is a collection of inventory slots and objects
	Inventory struct {
		Head      InventoryObject
		Back      InventoryObject
		RightHand InventoryObject
		LeftHand  InventoryObject
		Ring1     InventoryObject
		Ring2     InventoryObject
		Ring3     InventoryObject
	}
)
