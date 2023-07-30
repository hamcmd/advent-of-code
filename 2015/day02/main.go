package main

// https://adventofcode.com/2015/day/2

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

func main() {

	questionPart := 2
	content, err := utils.ReadStringFromFile("2015/day02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if questionPart == 1 {
		wrappingPaperRequired, err := calculateWrappingPaperRequired(&content)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(wrappingPaperRequired)
		return
	}

	ribbonRequired, err := calculateRibbonRequired(&content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ribbonRequired)
}

func calculateWrappingPaperRequired(stringDimensions *string) (int, error) {
	var totalSqFt int

	for _, value := range strings.Split(*stringDimensions, "\n") {
		var length, width, height int

		_, err := fmt.Sscanf(value, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			return 0, err
		}

		totalSqFt += 2 * length * width
		totalSqFt += 2 * width * height
		totalSqFt += 2 * height * length
		totalSqFt += utils.Min(
			length*width,
			width*height,
			height*length,
		)
	}

	return totalSqFt, nil
}

func calculateRibbonRequired(stringDimensions *string) (int, error) {
	var totalSqFtOfRibbon int

	for _, value := range strings.Split(*stringDimensions, "\n") {
		var length, width, height int

		_, err := fmt.Sscanf(value, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			return 0, err
		}

		volume := length * width * height
		totalSqFtOfRibbon += volume

		perimeters := []int{
			2 * (length + width),
			2 * (width + height),
			2 * (height + length),
		}

		perimeterOfSmallestSide := utils.Min(perimeters...)
		totalSqFtOfRibbon += perimeterOfSmallestSide
	}

	return totalSqFtOfRibbon, nil
}
