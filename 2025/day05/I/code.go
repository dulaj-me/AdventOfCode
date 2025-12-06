package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/viduranga/AdventOfCode/util"
)

type Range struct {
	min int
	max int
}

func (r *Range) Contains(n int) bool {
	return n >= r.min && n <= r.max
}

func parseRanges(s []string) []Range {
	ranges := make([]Range, 0, len(s))
	for _, str := range s {
		split := strings.Split(str, "-")

		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])

		ranges = append(ranges, Range{min: min, max: max})
	}

	return ranges
}

func parseIngredientIds(s []string) []int {
	ids := make([]int, 0, len(s))

	for _, str := range s {
		id, _ := strconv.Atoi(str)
		ids = append(ids, id)
	}

	return ids
}

func ingredientIsFresh(ranges []Range, ingredientId int) bool {
	for _, r := range ranges {
		if r.Contains(ingredientId) {
			return true
		}
	}
	return false
}

func filterFreshIngredients(ranges []Range, ingredientIds []int) []int {

	freshIngredients := make([]int, 0)

	for _, id := range ingredientIds {
		if ingredientIsFresh(ranges, id) {
			freshIngredients = append(freshIngredients, id)
		}
	}

	return freshIngredients

}

func main() {
	path := os.Args[1]
	group, err := util.FileToLineGroups(path)
	if err != nil {
		panic(err)
	}

	ranges := parseRanges(group[0])
	ingredientIds := parseIngredientIds(group[1])

	freshIngredients := filterFreshIngredients(ranges, ingredientIds)

	fmt.Println(len(freshIngredients))
}
