package main

//https://adventofcode.com/2015/day/6

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	questionPart := 2
	content, err := utils.ReadStringFromFile("2015/day06/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var answer int

	if questionPart == 1 {
		answer, err = handleLights(&content)
	} else {
		answer, err = handleBrightness(&content)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}

func handleLights(input *string) (int, error) {
	lightGrid := make([][]bool, 1000)

	for i := range lightGrid {
		lightGrid[i] = make([]bool, 1000)
	}

	lightCount := 0
	for _, line := range strings.Split(*input, "\n") {
		instruction := string(line)
		var x1, y1, x2, y2 int
		if strings.HasPrefix(instruction, "turn on") {
			_, err := fmt.Sscanf(instruction, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if !lightGrid[i][j] {
						lightGrid[i][j] = true
						lightCount++
					}
				}
			}
		}

		if strings.HasPrefix(instruction, "turn off") {
			_, err := fmt.Sscanf(instruction, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if lightGrid[i][j] {
						lightGrid[i][j] = false
						lightCount--
					}
				}
			}
		}

		if strings.HasPrefix(instruction, "toggle") {
			_, err := fmt.Sscanf(instruction, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lightGrid[i][j] = !lightGrid[i][j]

					if lightGrid[i][j] {
						lightCount++
					} else {
						lightCount--
					}

				}
			}
		}
	}

	return lightCount, nil
}

func handleBrightness(input *string) (int, error) {
	lightGrid := make([][]int, 1000)

	for i := range lightGrid {
		lightGrid[i] = make([]int, 1000)
	}

	lightBrightness := 0
	for _, line := range strings.Split(*input, "\n") {
		instruction := string(line)
		var x1, y1, x2, y2 int
		if strings.HasPrefix(instruction, "turn on") {
			_, err := fmt.Sscanf(instruction, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lightGrid[i][j]++
					lightBrightness++
				}
			}
		}

		if strings.HasPrefix(instruction, "turn off") {
			_, err := fmt.Sscanf(instruction, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if lightGrid[i][j] > 0 {
						lightGrid[i][j]--
						lightBrightness--
					}
				}
			}
		}

		if strings.HasPrefix(instruction, "toggle") {
			_, err := fmt.Sscanf(instruction, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)

			if err != nil {
				return 0, fmt.Errorf("wrong input format for instruction: %s", instruction)
			}

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lightGrid[i][j] += 2
					lightBrightness += 2
				}
			}
		}
	}

	return lightBrightness, nil
}
