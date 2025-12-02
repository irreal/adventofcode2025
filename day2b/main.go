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
			for window := 1; window <= l/2; window++ {
				if l%window != 0 {
					continue // Only lengths that divide evenly
				}
				pattern := numStr[:window]
				isRepeat := true
				for i := 0; i < l; i += window {
					if numStr[i:i+window] != pattern {
						isRepeat = false
						break
					}
				}
				if isRepeat {
					repeatingStrings = append(repeatingStrings, numStr)
					break
				}
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
	fmt.Printf("Sum is %d", sum)
}
