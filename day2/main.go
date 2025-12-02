package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	repeatingStrings := []string{}
	ranges := strings.Split(strings.TrimSpace(input), ",")
	for _, r := range ranges {
		parts := strings.Split(strings.TrimSpace(r), "-")
		if len(parts) != 2 {
			continue // skip invalid range
		}
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || start > end {
			continue // skip invalid range
		}
		for num := start; num <= end; num++ {
			numStr := fmt.Sprintf("%d", num)
			l := len(numStr)
			if l%2 != 0 {
				continue // Only even length can repeat exactly twice
			}
			half := l / 2
			// Only need one check: split in half and compare
			if numStr[:half] == numStr[half:] {
				repeatingStrings = append(repeatingStrings, numStr)
			}
		}
	}
	// Print all repeating strings found
	sum := 0
	for _, s := range repeatingStrings {
		fmt.Println(s)
		n, err := strconv.Atoi(s)
		if err == nil {
			sum += n
		}
	}
	fmt.Printf("Sum is %d\n", sum)
}
