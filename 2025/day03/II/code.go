package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/viduranga/AdventOfCode/util"
)

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

func getJoltage(bank []int, digitCount int) []int {

	endPoint := len(bank) - digitCount + 1
	masked := bank[:endPoint]
	maxJoltage := util.ArrayMax(masked)

	maxJoltageId := util.ArrayFindFirst(masked, maxJoltage)

	joltage := []int{maxJoltage}
	if digitCount == 1 {
		return joltage
	} else {

		return append(joltage, getJoltage(bank[maxJoltageId+1:], digitCount-1)...)
	}
}

func numArrayToNum(arr []int) int {

	result := 0

	for _, digit := range arr {
		result = result*10 + digit
	}

	return result
}

func calcJoltages(banks [][]int) []int {
	joltages := make([]int, 0)
	for _, bank := range banks {
		joltageArr := getJoltage(bank, 12)
		joltage := numArrayToNum(joltageArr)

		joltages = append(joltages, joltage)
	}

	return joltages

}

func main() {
	path := os.Args[1]
	lines, err := util.FileToLines(path)
	if err != nil {
		panic(err)
	}

	data, err := linesToArrays(lines)

	joltages := calcJoltages(data)

	fmt.Println(util.ArraySum(joltages))
}
