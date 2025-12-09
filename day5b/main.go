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

	for _, row := range rows {
		if !strings.Contains(row, "-") {
			break
		}

		parts := strings.Split(row, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	ranges = reduceRanges(ranges)

	countItems := 0

	for _, r := range ranges {
		countItems += r.End - r.Start + 1
	}

	// log.Printf("Ranges: %v\n", ranges)
	// log.Printf("Items: %v\n", items)
	log.Printf("Count Items: %d\n", countItems)
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
