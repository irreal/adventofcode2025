package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	arr := make([][]int, len(rows))
	outputArr := make([][]string, len(rows))

	for idx, row := range rows {
		arr[idx] = make([]int, len(row))
		outputArr[idx] = make([]string, len(row))
		for i, char := range row {
			outputArr[idx][i] = string(char)
			if char == '.' {
				arr[idx][i] = 0
			} else {
				arr[idx][i] = 1
			}
		}
	}
	count := 0
	for y, row := range arr {
		for x := range row {
			if arr[y][x] == 1 && getNumberOfOnesSurrounding(arr, x, y) < 4 {
				outputArr[y][x] = "X"
				count++
			}
			fmt.Print(outputArr[y][x])
		}
		fmt.Println()
	}
	fmt.Printf("Count: %d", count)
}

func getRelativeCoordinates(arr [][]int, x int, y int, deltaX int, deltaY int) int {
	newY := y + deltaY
	newX := x + deltaX

	if newY < 0 || newY >= len(arr) || newX < 0 || (len(arr) > 0 && newX >= len(arr[0])) {
		return 0
	}
	return arr[newY][newX]
}

func getSurroundingCoordinates(arr [][]int, x int, y int) []int {
	return []int{
		getRelativeCoordinates(arr, x, y, -1, -1),
		getRelativeCoordinates(arr, x, y, 0, -1),
		getRelativeCoordinates(arr, x, y, 1, -1),
		getRelativeCoordinates(arr, x, y, -1, 0),
		getRelativeCoordinates(arr, x, y, 1, 0),
		getRelativeCoordinates(arr, x, y, -1, 1),
		getRelativeCoordinates(arr, x, y, 0, 1),
		getRelativeCoordinates(arr, x, y, 1, 1),
	}
}

func getNumberOfOnesSurrounding(arr [][]int, x int, y int) int {
	surrounding := getSurroundingCoordinates(arr, x, y)
	count := 0
	for _, cell := range surrounding {
		if cell == 1 {
			count++
		}
	}
	return count
}
