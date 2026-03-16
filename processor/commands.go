package processor

import (
	"strconv"
	"strings"
)

func applyCommands(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {

		if strings.HasPrefix(tokens[i], "(") && strings.HasSuffix(tokens[i], ")") {

			command := tokens[i][1 : len(tokens[i])-1]
			parts := strings.Split(command, ",")

			action := strings.TrimSpace(parts[0])
			count := 1

			if len(parts) == 2 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					count = n
				}
			}

			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				switch action {
				case "up":
					tokens[j] = strings.ToUpper(tokens[j])
				case "low":
					tokens[j] = strings.ToLower(tokens[j])
				case "cap":
					tokens[j] = capitalize(tokens[j])
				case "hex":
					tokens[j] = hexToDecimal(tokens[j])
				case "bin":
					tokens[j] = binToDecimal(tokens[j])
				}
			}

			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}
	return tokens
}
