package things

type (
	// ItemCollection is a collection of items
	ItemCollection []*Item
)

// Match returns matching items and strengths from the collection
func (col ItemCollection) Match(text string) *Item {
	var foundItem *Item
	var foundStrength = 0
	for _, itm := range col {
		if str := itm.Match(text); str > foundStrength {
			foundItem = itm
		}
	}
	return foundItem
}
