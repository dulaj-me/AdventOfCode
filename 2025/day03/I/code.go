package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/viduranga/AdventOfCode/util"
)

type Maxs struct {
	left  []int
	right []int
}

func linesToArrays(lines []string) ([][]int, error) {
	var result [][]int

	for _, line := range lines {
		lineInts := make([]int, 0, len(line))

		for _, char := range line {
			intVal, _ := strconv.Atoi(string(char))
			lineInts = append(lineInts, intVal)
		}

		result = append(result, lineInts)
	}
	return result, nil
}

func calcMaxs(lines [][]int) ([]Maxs, error) {

	maxsPerLine := make([]Maxs, 0)

	for _, line := range lines {
		lineLen := len(line)
		left := make([]int, lineLen-1)
		right := make([]int, lineLen-1)

		left[0] = line[0]
		right[lineLen-2] = line[lineLen-1]

		for i := 1; i < len(line)-1; i++ {
			j := lineLen - 2 - i

			left[i] = util.MaxInt(left[i-1], line[i])
			right[j] = util.MaxInt(right[j+1], line[j+1])
		}

		maxs := Maxs{
			left:  left,
			right: right,
		}

		maxsPerLine = append(maxsPerLine, maxs)
	}
	return maxsPerLine, nil
}

func calcLargestJoltages(maxs []Maxs) []int {

	largestJoltages := make([]int, 0)

	for _, max := range maxs {
		maxJoltage := 0
		for i := 0; i < len(max.left); i++ {
			maxJoltage = util.MaxInt(maxJoltage, max.left[i]*10+max.right[i])
		}

		largestJoltages = append(largestJoltages, maxJoltage)
	}

	return largestJoltages

}

func main() {
	path := os.Args[1]
	lines, err := util.FileToLines(path)
	if err != nil {
		panic(err)
	}

	data, err := linesToArrays(lines)

	maxs, err := calcMaxs(data)

	largestJoltages := calcLargestJoltages(maxs)

	fmt.Println(util.ArraySum(largestJoltages))
}
