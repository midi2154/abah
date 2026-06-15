package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	graph := make(map[rune][]string)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading file")
	}

	if len(data) == 0 {
		return nil, errors.New("empty file")
	}

	lines := strings.Split(string(data), "\n")

	for i := ' '; i <= '~'; i++ {
		start := int(i-32) * 9

		if start+8 >= len(lines) {
			return nil, errors.New("incomplete banner file")
		}

		graph[i] = lines[start+1 : start+9]
	}

	return graph, nil
}
