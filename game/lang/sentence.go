package lang

import (
	"fmt"
	"strings"
)

// MakeSentance constructs a sentence from a string array.
func MakeSentence(words []string) string {
	return fmt.Sprintf("%s and %s",
		strings.Join(words[:len(words)-1], ", "),
		words[len(words)-1])
}
