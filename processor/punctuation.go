package processor

func rebuild(tokens []string) string {
	result := ""

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if i > 0 && !isRightAttached(token) {
			result += " "
		}

		result += token
	}

	return result
}

func isRightAttached(token string) bool {
	switch token {
	case ".", ",", "!", "?", ":", ";", "...":
		return true
	}
	return false
}
