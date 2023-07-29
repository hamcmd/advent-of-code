package utils

import (
	"os"
	"strings"
)

func ReadStringFromFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func ReadLinesFromFileInStringArray(path string) ([]string, error) {
	content, err := ReadStringFromFile(path)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(content, "\n")

	return lines, err
}

func Min(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	minValue := numbers[0]

	for _, num := range numbers {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}
