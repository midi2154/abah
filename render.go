package main

import "strings"

func Render(s string, banner map[rune][]string) string {
	lines := strings.Split(s, "\\n")

	var result string

	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					result += art[row]
				}
			}
			result += "\n"
		}
	}

	return result
}
