package main

import "strings"

func Render(input string, banner map[rune][]string) string {
	textLines := strings.Split(input, "\\n")

	output := ""
	allLinesEmpty := true

	for _, line := range textLines {
		if line == "" {
			output += "\n"
			continue
		}

		allLinesEmpty = false

		for row := 0; row < 8; row++ {
			for _, char := range line {
				if artRows, exists := banner[char]; exists {
					output += artRows[row]
				}
			}
			output += "\n"
		}
	}

	if allLinesEmpty {
		return strings.Repeat("\n", len(textLines)-1)
	}

	return output
}
