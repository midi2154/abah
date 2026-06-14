package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	graph := map[rune][]string{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading file")
	}

	if len(data) == 0 {
		return nil, errors.New("empty file")
	}

	result := strings.Split(string(data), "\n")

	if len(result) < 855 {
		return nil, errors.New("incomplete banner file")
	}

	for i := ' '; i <= '~'; i++ {
		start := int(i-32) * 9
		graph[i] = result[start+1 : start+9]
	}

	return graph, nil
}
