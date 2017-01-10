package things

import (
	"rift/game/db"
	"strings"

	"regexp"

	logging "github.com/op/go-logging"

	"gopkg.in/mgo.v2/bson"
)

const (
	dbCollectionName = "items"
)

var (
	log = logging.MustGetLogger("items")
)

type (
	// Item is a instantiated item form that exists somewhere in the world
	Item struct {
		ID       bson.ObjectId
		Form     ItemForm
		Items    ItemCollection
		Location bson.ObjectId
	}
)

// Match checks a text string agianst the item  it returns a relevence score
func (item *Item) Match(text string) int {
	strength := 0
	words := strings.Split(text, " ")
	nounWord := words[len(words)-1]
	adjectiveWord := ""
	nounPattern := []byte(item.Form.Noun)
	adjPattern := []byte(item.Form.PrimaryAdjective)
	if len(words) > 1 {
		adjectiveWord = words[len(words)-2]
	}
	if m, _ := regexp.Match(nounWord, nounPattern); m {
		strength += len(nounWord)
		if len(words) > 1 {
			if m, _ := regexp.Match(adjectiveWord, adjPattern); m {
				strength += len(adjectiveWord)
			}
		}
		return strength
	}
	return 0
}

// MoveItem changes the location of an item
func (item *Item) MoveItem(newLocationID bson.ObjectId) {
	item.Location = newLocationID
}

// ContainItem places a given item inside the item
func (item *Item) ContainItem(containedItem *Item) bool {
	item.Items = append(item.Items, containedItem)
	containedItem.MoveItem(item.ID)
	return true
}

// NewItem instatiates and item from an item form
func NewItem(form *ItemForm) *Item {
	item := &Item{
		ID:       bson.NewObjectId(),
		Form:     *form,
		Location: bson.NewObjectId(),
	}
	return item
}

// LoadItem retrieves an item from the database
func LoadItem(id bson.ObjectId) *Item {
	ses, c := db.GetCollection(dbCollectionName)
	defer ses.Close()
	var item Item
	err := c.Find(bson.M{"_id": id}).One(&item)
	if err != nil {
		log.Critical(err)
		return nil
	}
	return &item
}

// Save saves an item in the database
func (item *Item) Save() error {
	ses, c := db.GetCollection(dbCollectionName)
	defer ses.Close()
	_, err := c.Upsert(bson.M{"_id": item.ID}, item)
	return err
}
