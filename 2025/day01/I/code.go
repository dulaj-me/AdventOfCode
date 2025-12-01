package main

import (
	"fmt"
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
	var current = 50
	var password = 0

	for _, turn := range turns {
		current += turn
		current = mod(current, 100)

		if current == 0 {
			password += 1
		}
	}

	return password

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
