package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Range struct {
	Start int
	End   int // Inclusive
}

func main() {
	rows := strings.Split(input, "\n")
	ranges := make([]Range, 0, len(rows))
	items := make([]int, 0, len(rows))

	mode := "ranges"

	for _, row := range rows {
		if mode == "ranges" {
			if !strings.Contains(row, "-") {
				mode = "items"
				continue
			}

			parts := strings.Split(row, "-")
			start, err := strconv.Atoi(parts[0])
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			ranges = append(ranges, Range{Start: start, End: end})
		} else {
			item, err := strconv.Atoi(row)
			if err != nil {
				continue
			}
			items = append(items, item)
		}
	}

	ranges = reduceRanges(ranges)

	countFresh := 0

	for _, item := range items {
		if contains(ranges, item) {
			countFresh++
		}
	}

	// log.Printf("Ranges: %v\n", ranges)
	// log.Printf("Items: %v\n", items)
	log.Printf("Count Fresh: %d\n", countFresh)
}

func contains(ranges []Range, item int) bool {
	for _, r := range ranges {
		if item < r.Start {
			return false
		}
		if item <= r.End {
			return true
		}
	}
	return false
}

func reduceRanges(ranges []Range) []Range {

	// Sort by Start asc, then End asc
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].Start != ranges[j].Start {
			return ranges[i].Start < ranges[j].Start
		}
		return ranges[i].End < ranges[j].End
	})

	merged := make([]Range, 0, len(ranges))
	merged = append(merged, ranges[0])

	for i := 1; i < len(ranges); i++ {
		curr := ranges[i]
		last := &merged[len(merged)-1]

		limit := last.End + 1

		if curr.Start <= limit {
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged

}
