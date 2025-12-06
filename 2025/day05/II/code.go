package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/viduranga/AdventOfCode/util"
)

type Range struct {
	min uint64
	max uint64
}

func (a *Range) Merge(b Range) (merged *Range, ok bool) {
	if a.min <= b.max && b.min <= a.max {
		return &Range{min: util.Min(a.min, b.min), max: util.Max(a.max, b.max)}, true
	}

	return nil, false
}

func mergeRanges(ranges []Range) []Range {
	rangeMap := make(map[int]Range)
	for i, r := range ranges {
		rangeMap[i] = r
	}

	for {
		mergedAny := false

		// It's ok this is random order, we just need to keep merging until no more merges are possible
		for i, r1 := range rangeMap {
			for j, r2 := range rangeMap {
				if i == j {
					continue
				}
				merged, ok := r1.Merge(r2)
				if ok {
					mergedAny = true
					rangeMap[j] = *merged
					delete(rangeMap, i)
				}
			}
		}

		if !mergedAny {
			break
		}
	}

	newRanges := make([]Range, 0, len(rangeMap))

	for _, r := range rangeMap {
		newRanges = append(newRanges, r)
	}
	return newRanges
}

func parseRanges(s []string) []Range {
	ranges := make([]Range, 0, len(s))
	for _, str := range s {
		split := strings.Split(str, "-")

		min, _ := strconv.ParseUint(split[0], 10, 64)
		max, _ := strconv.ParseUint(split[1], 10, 64)

		ranges = append(ranges, Range{min: min, max: max})
	}

	return ranges
}

func idCount(ranges []Range) uint64 {

	count := uint64(0)

	for _, r := range ranges {
		count += r.max - r.min + 1
	}

	return count

}

func main() {
	path := os.Args[1]
	group, err := util.FileToLineGroups(path)
	if err != nil {
		panic(err)
	}

	ranges := parseRanges(group[0])

	margedRanges := mergeRanges(ranges)

	fmt.Println("Merged Ranges:", margedRanges)

	fmt.Println(idCount(margedRanges))
}
