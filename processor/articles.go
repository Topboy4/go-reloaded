package processor

import "strings"

func fixArticles(tokens []string) {
	for i := 0; i < len(tokens)-1; i++ {
		if strings.ToLower(tokens[i]) == "a" {
			next := strings.ToLower(tokens[i+1])
			if startsWithVowelOrH(next) {
				if tokens[i] == "A" {
					tokens[i] = "An"
				} else {
					tokens[i] = "an"
				}
			}
		}
	}
}

func startsWithVowelOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	return strings.ContainsRune("aeiouh", rune(word[0]))
}
