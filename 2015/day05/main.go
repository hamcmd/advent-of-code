package main

//https://adventofcode.com/2015/day/5

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	questionPart := 2
	content, err := utils.ReadStringFromFile("2015/day05/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var niceWords int

	if questionPart == 1 {
		niceWords = countNiceWordsPart1(&content)
	} else {
		niceWords = countNiceWordsPart2(&content)
	}

	fmt.Println(niceWords)
}

func countNiceWordsPart1(input *string) int {
	nice := 0
	restrictedWords := regexp.MustCompile("ab|cd|pq|xy")
	vowels := regexp.MustCompile("a|e|i|o|u")

	for _, word := range strings.Split(*input, "\n") {
		str := string(word)
		if restrictedWords.FindStringSubmatch(str) != nil {
			continue
		}

		vowelMatches := vowels.FindAllStringSubmatch(str, -1)

		if len(vowelMatches) < 3 {
			continue
		}

		for index := 0; index < len(str)-1; index++ {
			if str[index] == str[index+1] {
				nice++
				break
			}
		}
	}

	return nice
}

func countNiceWordsPart2(input *string) int {
	nice := 0

	for _, line := range strings.Split(*input, "\n") {
		word := string(line)
		var passesRule1 bool

		for i := 0; i < len(word)-2; i++ {
			toMatch := word[i : i+2]
			for j := i + 2; j < len(word)-1; j++ {
				if word[j:j+2] == toMatch {
					passesRule1 = true
					break
				}
			}
		}

		for index := 0; index < len(word)-2; index++ {
			if word[index] == word[index+2] && passesRule1 {
				nice++
				break
			}
		}
	}

	return nice
}
