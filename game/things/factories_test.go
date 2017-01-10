package things

func shortSwordFormFactory() *ItemForm {
	form := NewItemForm()
	form.Noun = "sword"
	form.FullName = "short sword"
	form.PrimaryAdjective = "short"
	return form
}

func sheathFormFactory() *ItemForm {
	form := NewItemForm()
	form.Noun = "sheath"
	form.PrimaryAdjective = "leather"
	form.FullName = "leather sheath"
	form.Container = true
	return form
}

func shortSwordFactory() *Item {
	form := shortSwordFormFactory()
	item := NewItem(form)
	return item
}

func sheathFactory() *Item {
	form := sheathFormFactory()
	item := NewItem(form)
	return item
}
