package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/viduranga/AdventOfCode/2023/util"
)

func ranges(line string) ([][2]string, error) {
	re := regexp.MustCompile(`(\d+)-(\d+)$`)

	var result [][2]string
	data := strings.Split(line, ",")

	for _, token := range data {
		match := re.FindStringSubmatch(token)

		if len(match) == 0 {
			return nil, fmt.Errorf("match failed in %s", line)
		}
		first, second := match[1], match[2]

		result = append(result, [2]string{first, second})
	}

	return result, nil
}

func sillyPatternsInRange(first, second string) []int {
	firstInt, _ := strconv.Atoi(first)
	secondInt, _ := strconv.Atoi(second)
	firstLength := len(first)
	secondLength := len(second)

	var firstHalfLen, secondHalfLen int
	if firstLength == secondLength {
		firstHalfLen = int(math.Ceil(float64(firstLength) / 2))
		secondHalfLen = int(math.Ceil(float64(secondLength) / 2))
	} else {
		firstHalfLen = int(math.Floor(float64(firstLength) / 2))
		secondHalfLen = int(math.Ceil(float64(secondLength) / 2))
	}

	firstFront, _ := strconv.Atoi(first[:firstHalfLen])
	secondFront, _ := strconv.Atoi(second[:secondHalfLen])

	var sillyPatterns []int

	for i := firstFront; i <= secondFront; i++ {
		candidateStr := fmt.Sprintf("%d%d", i, i)
		candidate, _ := strconv.Atoi(candidateStr)

		if candidate >= firstInt && candidate <= secondInt {
			sillyPatterns = append(sillyPatterns, candidate)
		}
	}

	return sillyPatterns
}

func getSillyPatterns(rangeData [][2]string) []int {
	var sillyPatterns []int
	for _, r := range rangeData {
		sillyPatternsInRange := sillyPatternsInRange(r[0], r[1])

		sillyPatterns = append(sillyPatterns, sillyPatternsInRange...)

	}

	return sillyPatterns
}

func main() {
	path := os.Args[1]
	lines, err := util.FileToLines(path)
	if err != nil {
		panic(err)
	}

	rangeData, err := ranges(lines[0])
	if err != nil {
		panic(err)
	}

	sillyPatterns := getSillyPatterns(rangeData)

	fmt.Println(util.ArraySum(sillyPatterns))
}
