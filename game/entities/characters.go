package entities

import (
	"rift/game/db"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Character is a basic game character either an NPC or player controlled
	Character struct {
		ID        bson.ObjectId `bson:"_id"`
		Name      string
		Inventory Inventory
	}
	// CharacterCollection is a collection of character objects
	CharacterCollection []*Character
)

// Names returns an array of character names
func (cc CharacterCollection) Names() []string {
	names := []string{}
	for _, c := range cc {
		names = append(names, c.Name)
	}
	return names
}

// ShortNames returns a slice of displayable short names for the characters
func (cc CharacterCollection) ShortNames() []string {
	names := []string{}
	for _, c := range cc {
		names = append(names, c.Name)
	}
	return names
}

// LoadCharacter loads a character from the database
func LoadCharacter(id bson.ObjectId) (*Character, error) {
	c := db.GetCollection("characters")
	var char Character
	err := c.Find(bson.M{"_id": id}).One(&char)
	return &char, err
}

//NewCharacter constructs a new character
func NewCharacter() *Character {
	return &Character{
		ID: bson.NewObjectId(),
	}
}
