package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/viduranga/AdventOfCode/2023/util"
)

func turns(data []string) ([]int, error) {
	re := regexp.MustCompile(`([L-R])(\d+)$`)

	var result []int

	for _, line := range data {
		match := re.FindStringSubmatch(line)

		if len(match) == 0 {
			return nil, fmt.Errorf("match failed in %s", line)
		}
		direction, count := match[1], match[2]

		digit, err := strconv.Atoi(count)
		if err != nil {
			return nil, err
		}

		if direction == "L" {
			digit = -digit
		}

		result = append(result, digit)
	}

	return result, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func password(turns []int) int {
	current := 50
	zeroCrossings := 0
	zeroLandings := 0

	for _, turn := range turns {
		next := current + turn

		crossings := int(math.Abs(float64(next / 100)))

		if current != 0 && next < 0 { // we crossed zero going negative
			crossings += 1
		}

		next = mod(next, 100)

		if next == 0 {
			if crossings > 0 {
				crossings -= 1 // reduce one because we landed on zero
			}
			zeroLandings += 1
		}
		zeroCrossings += crossings

		current = next

	}

	return zeroCrossings + zeroLandings

}

func main() {
	path := os.Args[1]
	lines, err := util.FileToLines(path)
	if err != nil {
		panic(err)
	}

	turns, err := turns(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(password(turns))
}
