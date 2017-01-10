package things

import "testing"

import "rift/game/db"

var shortSwordForm *ItemForm

func TestNewItem(t *testing.T) {
	shortSwordForm := NewItemForm()
	shortSwordForm.Noun = "sword"
	shortSwordForm.PrimaryAdjective = "short"
	shortSword := NewItem(shortSwordForm)
	if shortSword == nil {
		t.Error("Item was not instatiated from NewItem")
	}
}

func TestSave(t *testing.T) {
	db.GetCollection(dbCollectionName).DropCollection()
	sword := shortSwordFactory()
	sword.Save()
	col := db.GetCollection(dbCollectionName)
	if count, _ := col.Count(); count != 1 {
		t.Errorf("Record not saved to collection, count is %d", count)
	}
}

func TestContainItem(t *testing.T) {
	shortSword := shortSwordFactory()
	sheath := sheathFactory()
	sheath.ContainItem(shortSword)
	if shortSword.Location != sheath.ID {
		t.Error("item location is not the container location")
	}
}

func TestLoadItem(t *testing.T) {
	sword := shortSwordFactory()
	if err := sword.Save(); err != nil {
		t.Error(err)
	}

	item := LoadItem(sword.ID)
	if item == nil {
		t.Error("item not loaded")
	}

	if item.Form.Noun != sword.Form.Noun {
		t.Errorf("item not loaded properly, noun should be %s, got %s",
			sword.Form.Noun, item.Form.Noun)
	}
}

func TestMatch(t *testing.T) {
	shortSword := shortSwordFactory()
	if str := shortSword.Match("sh sw"); str != 4 {
		t.Errorf("sh sw should return 4, got %d", str)
	}

	if str := shortSword.Match("sword"); str != 5 {
		t.Errorf("sh sw should return 4, got %d", str)
	}

	if str := shortSword.Match("short sword"); str != 10 {
		t.Errorf("sh sw should return 4, got %d", str)
	}

	if str := shortSword.Match("words"); str != 0 {
		t.Errorf("sh sw should return 0, got %d", str)
	}

}
