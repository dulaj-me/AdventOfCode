package main

import (
	"fmt"
	"os"

	"github.com/viduranga/AdventOfCode/util"
)

type Maxs struct {
	left  []int
	right []int
}

func surroundingRollsCount(grid [][]rune, i, j int) int {

	surroundingRolls := 0

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			_i := i + di
			_j := j + dj

			if _i < 0 || _i >= len(grid) || _j < 0 || _j >= len(grid[0]) {
				continue
			}

			if grid[_i][_j] == '@' {
				surroundingRolls++
			}
		}
	}

	return surroundingRolls
}

func accessibleRolls(grid [][]rune) int {
	accessible := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == '.' {
				continue
			} else if cell == '@' {
				surrounding := surroundingRollsCount(grid, i, j)

				if surrounding < 4 {
					accessible++
				}
			}

		}
	}

	return accessible
}

func main() {
	path := os.Args[1]
	grid, err := util.FileToGrid(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(accessibleRolls(grid))
}
