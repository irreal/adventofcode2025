package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed day1input.txt
var input string

func main() {
	maxDial := 100

	dial := 50

	zeroCount := 0

	for _, instruction := range strings.Split(input, "\n") {
		direction := instruction[0]
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic("invalid instruction: " + instruction)
		}

		switch direction {
		case 'L':
			dial -= steps
		case 'R':
			dial += steps
		}

		dial := dial % maxDial
		if dial == 0 {
			zeroCount++
		}
	}

	log.Printf("zeroCount: %d", zeroCount)
}
