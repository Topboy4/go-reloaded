package processor

import "strings"

func Process(text string) string {

	// Remove UTF-8 BOM if present
	text = strings.TrimPrefix(text, "\uFEFF")

	tokens := tokenize(text)

	tokens = applyCommands(tokens)

	fixArticles(tokens)

	return rebuild(tokens)
}
