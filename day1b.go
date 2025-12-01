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

		zeroCount += steps / maxDial
		steps %= maxDial
		startedAtZero := dial == 0

		switch direction {
		case 'L':
			dial -= steps
			if dial <= 0 {
				if !startedAtZero {
					zeroCount++
				}
				if dial < 0 {
					dial += maxDial
				}
			}
		case 'R':
			dial += steps
			if dial >= maxDial {
				zeroCount++
				dial -= maxDial
			}
		}

	}

	log.Printf("zeroCount: %d", zeroCount)
}
