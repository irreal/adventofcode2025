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
		maxStrength := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				strength, err := strconv.Atoi(bank[i:i+1] + bank[j:j+1])
				if err != nil {
					continue
				}
				if strength > maxStrength {
					maxStrength = strength
				}
			}
		}
		arr[idx] = maxStrength
	}
	total := 0
	for _, strength := range arr {
		total += strength
	}
	// fmt.Printf("Arr: %v", arr)
	fmt.Printf("Total: %d", total)
}
