package cleaninput

import (
	"strings"
)

func CleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
