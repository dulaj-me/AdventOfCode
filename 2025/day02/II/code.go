package main

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
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

		firstLen := len(first)
		secondLen := len(second)

		if firstLen == secondLen {

			result = append(result, [2]string{first, second})
		} else if firstLen+1 == secondLen {
			midpoint := strings.Repeat("9", firstLen)
			midpointPlusOne := "1" + strings.Repeat("0", firstLen)
			result = append(result, [2]string{first, midpoint})
			result = append(result, [2]string{midpointPlusOne, second})
		} else {
			return nil, fmt.Errorf("invalid range lengths in %s-%s", first, second)
		}

	}

	return result, nil
}

func sillyPatternsInRange(first, second string) []int {
	firstInt, _ := strconv.Atoi(first)
	secondInt, _ := strconv.Atoi(second)
	numLen := len(first)

	sillyPatterns := make(map[int]bool)

	for i := 1; i <= numLen/2; i++ {
		if numLen%i != 0 {
			continue
		}

		repeatCount := numLen / i

		firstFront, _ := strconv.Atoi(first[:i])
		secondFront, _ := strconv.Atoi(second[:i])

		for j := firstFront; j <= secondFront; j++ {
			repeatedStr := strings.Repeat(strconv.Itoa(j), repeatCount)

			repeatedInt, _ := strconv.Atoi(repeatedStr)

			if repeatedInt >= firstInt && repeatedInt <= secondInt {
				if !sillyPatterns[repeatedInt] {
					sillyPatterns[repeatedInt] = true
				}
			}
		}
	}

	return slices.Collect(maps.Keys(sillyPatterns))
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
