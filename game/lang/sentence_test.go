package lang

import (
	"testing"
)

func TestMakeSentence(t *testing.T) {
	items := []string{"one", "two", "three"}
	sentence := MakeSentence(items)
	if sentence != "one, two and three" {
		t.Error("Sentence was not constructed correctly")
	}
}
