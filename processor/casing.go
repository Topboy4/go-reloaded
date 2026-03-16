package processor

import "strings"

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}
