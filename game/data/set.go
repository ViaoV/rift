package collections

type (
	// Flags are a set collection for string based flags
	Flags []string
)

// Contains returns true if the flag is contained inside the collection
func (f Flags) Contains(flag string) bool {
	for _, i := range f {
		if i == flag {
			return true
	}
	return false
}
