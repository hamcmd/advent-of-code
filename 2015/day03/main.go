package main

//https://adventofcode.com/2015/day/3

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	questionPart := 2
	content, err := utils.ReadStringFromFile("2015/day03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var numberOfHouses int
	if questionPart == 1 {
		numberOfHouses = housesVisitedBySanta(content)
	} else {
		numberOfHouses = housesVisitedBySantaAndBot(content)
	}

	fmt.Println(numberOfHouses)
}

/////// current coordinates = (x, y)
//////  direction to move = (a, b)
//////  coordinates after moving or coordinates visited next = (x + a, y + b)
//////  number of unique coordinates visited = number of houses visisted

func housesVisitedBySanta(input string) int {
	directions := getCoordinateMapping()
	coordinatesVisted := map[[2]int]int{{}: 1}
	currentCoordinates := [2]int{0, 0}

	for _, val := range strings.Split(input, "") {
		directionToGo := directions[string(val)]
		nextCoorindates := [2]int{
			directionToGo[0] + currentCoordinates[0],
			directionToGo[1] + currentCoordinates[1],
		}

		currentCoordinates = nextCoorindates
		coordinatesVisted[currentCoordinates]++
	}

	return len(coordinatesVisted)
}

//// since two santas are moving together from the same starting point,
//// santa would visit coordinates at all odd positions
//// bot will be going to all the coordinates at even position

func housesVisitedBySantaAndBot(input string) int {
	directions := getCoordinateMapping()
	coordinatesVisted := map[[2]int]int{{}: 2}
	currentSantaCoordinates := [2]int{0, 0}
	currentBotCoordinates := [2]int{0, 0}

	for index, val := range strings.Split(input, "") {
		directionToGo := directions[string(val)]
		if index%2 == 0 {
			nextCoorindates := [2]int{
				directionToGo[0] + currentSantaCoordinates[0],
				directionToGo[1] + currentSantaCoordinates[1],
			}

			currentSantaCoordinates = nextCoorindates
			coordinatesVisted[currentSantaCoordinates]++
		} else {
			nextCoorindates := [2]int{
				directionToGo[0] + currentBotCoordinates[0],
				directionToGo[1] + currentBotCoordinates[1],
			}

			currentBotCoordinates = nextCoorindates
			coordinatesVisted[currentBotCoordinates]++
		}
	}

	return len(coordinatesVisted)
}

func getCoordinateMapping() map[string][2]int {
	return map[string][2]int{
		"^": {0, 1},
		"v": {0, -1},
		">": {1, 0},
		"<": {-1, 0},
	}
}
