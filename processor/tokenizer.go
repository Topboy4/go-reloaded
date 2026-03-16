package processor

import "unicode"

func tokenize(text string) []string {
	var tokens []string
	var current []rune

	for i := 0; i < len(text); {
		r := rune(text[i])

		// Handle ellipsis
		if i+2 < len(text) && text[i:i+3] == "..." {
			flushCurrent(&tokens, &current)
			tokens = append(tokens, "...")
			i += 3
			continue
		}

		// Handle commands ( ... )
		if r == '(' {
			flushCurrent(&tokens, &current)
			start := i
			for i < len(text) && text[i] != ')' {
				i++
			}
			i++
			tokens = append(tokens, text[start:i])
			continue
		}

		// Handle punctuation
		if isPunctuation(r) {
			flushCurrent(&tokens, &current)
			tokens = append(tokens, string(r))
			i++
			continue
		}

		// Handle space
		if unicode.IsSpace(r) {
			flushCurrent(&tokens, &current)
			i++
			continue
		}

		current = append(current, r)
		i++
	}

	flushCurrent(&tokens, &current)
	return tokens
}

func flushCurrent(tokens *[]string, current *[]rune) {
	if len(*current) > 0 {
		*tokens = append(*tokens, string(*current))
		*current = []rune{}
	}
}

func isPunctuation(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';', '\'':
		return true
	}
	return false
}
