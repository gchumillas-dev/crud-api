package utils

// InArray checkes that item is included in items.
func InArray(item string, items []string) bool {
	for index := range items {
		if items[index] == item {
			return true
		}
	}

	return false
}

// Unpack destructuring.
func Unpack(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}
