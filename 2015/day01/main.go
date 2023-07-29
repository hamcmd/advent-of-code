package main

//https://adventofcode.com/2015/day/1

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func main() {
	input, err := utils.ReadStringFromFile("2015/day01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	level := getLevel(input, 2)
	fmt.Println(level)
}

func getLevel(input string, questionPart int) int {
	var floor int
	for index, char := range input {
		if string(char) == "(" {
			floor++
		} else {
			floor--
		}

		if questionPart == 2 && floor == -1 {
			return index + 1
		}
	}

	return floor
}
