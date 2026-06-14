package main

import "strings"

func Render(s string, banner map[rune][]string) string {
	lines := strings.Split(s, "\\n")
	result := ""
	emptyString := true
	for _, i := range lines {
		if i == "" {
			result += "\n"
			continue
		}
		if i != "" {
			emptyString = false
		}
		for row := 0; row < 8; row++ {
			for _, ch := range i {
				if art, ok := banner[ch]; ok {
					result += art[row]
				}
			}
			result += "\n"
		}
	}
	if emptyString == true {
		return strings.Repeat("\n", len(lines)-1)
	}
	return result
}
