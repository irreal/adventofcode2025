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
	banks := strings.Split(input, "\n")
	arr := make([]int, len(banks))

	for idx, bank := range banks {
		digits := make([]int, len(bank))
		for i, n := range strings.Split(bank, "") {
			digits[i], _ = strconv.Atoi(n)
		}

		toKeep := 12
		if len(digits) < toKeep {
			toKeep = len(digits)
		}

		result := make([]int, 0, toKeep)
		start := 0

		for len(result) < toKeep {
			needed := toKeep - len(result)
			end := len(digits) - needed + 1

			maxIdx := start
			for i := start; i < end; i++ {
				if digits[i] > digits[maxIdx] {
					maxIdx = i
				}
			}

			result = append(result, digits[maxIdx])
			start = maxIdx + 1
		}

		stringPicked := make([]string, len(result))
		for i, n := range result {
			stringPicked[i] = strconv.Itoa(n)
		}
		strength, err := strconv.Atoi(strings.Join(stringPicked, ""))
		if err != nil {
			continue
		}
		arr[idx] = strength
	}
	total := 0
	for _, strength := range arr {
		total += strength
	}
	// fmt.Printf("Arr: %v", arr)
	fmt.Printf("Total: %d", total)
}
