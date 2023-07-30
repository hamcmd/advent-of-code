package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	questionPart := 2
	input := "ckczppom"

	number, err := findSmallestValidNumber(input, questionPart)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(number)
}

func findSmallestValidNumber(input string, questionPart int) (int, error) {
	maxInt := math.MaxInt64
	prefixCount := 6

	if questionPart == 1 {
		prefixCount = 5
	}

	for number := 0; number < maxInt; number++ {
		concatenatedInput := fmt.Sprint(input, number)
		hash := getMD5Hash(concatenatedInput)

		if strings.HasPrefix(hash[0:prefixCount], strings.Repeat("0", prefixCount)) {
			return number, nil
		}
	}

	return 0, fmt.Errorf("hash not found")
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
