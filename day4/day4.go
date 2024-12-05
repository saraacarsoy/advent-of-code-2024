package main

import (
	utils "advent-of-code-2024/inputs"
	"log"
)

func main() {
	input, err := utils.ReadAs2dArray("day4")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	part1(input)
}

func part1(input [][]string) {
	total := 0
	aTotal := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			char := input[row][col]

			if char == "X" {
				total += checkNeighbors(input, row, col)
			}

			if char == "A" {
				aTotal += checkADiagonals(input, row, col)
			}
		}
	}
	print(" Part 1 Total:  ", total)
	print(" Part 2 Total:  ", aTotal)
}

func checkNeighbors(grid [][]string, row, col int) int {
	count := 0
	count += checkHorizontal(grid, row, col)
	count += checkVertical(grid, row, col)
	count += checkDiagonals(grid, row, col)

	return count
}

func checkHorizontal(grid [][]string, row, col int) int {
	count := 0
	if col+3 < len(grid[row]) && grid[row][col+1] == "M" && grid[row][col+2] == "A" && grid[row][col+3] == "S" {
		count += 1
	}

	if col-3 >= 0 && grid[row][col-1] == "M" && grid[row][col-2] == "A" && grid[row][col-3] == "S" {
		count += 1
	}

	return count
}

func checkVertical(grid [][]string, row, col int) int {
	count := 0
	if row+3 < len(grid) && grid[row+1][col] == "M" && grid[row+2][col] == "A" && grid[row+3][col] == "S" {
		count += 1
	}

	if row-3 >= 0 && grid[row-1][col] == "M" && grid[row-2][col] == "A" && grid[row-3][col] == "S" {
		count += 1
	}

	return count
}

func checkDiagonals(grid [][]string, row, col int) int {
	count := 0
	// top-left to bottom-right
	if row+3 < len(grid) && col+3 < len(grid[row]) &&
		grid[row+1][col+1] == "M" &&
		grid[row+2][col+2] == "A" &&
		grid[row+3][col+3] == "S" {
		count += 1
	}

	// bottom-left to top-right
	if row-3 >= 0 && col+3 < len(grid[row]) &&
		grid[row-1][col+1] == "M" &&
		grid[row-2][col+2] == "A" &&
		grid[row-3][col+3] == "S" {
		count += 1
	}

	// top-right to bottom-left
	if row+3 < len(grid) && col-3 >= 0 &&
		grid[row+1][col-1] == "M" &&
		grid[row+2][col-2] == "A" &&
		grid[row+3][col-3] == "S" {
		count += 1
	}

	// bottom-right to top-left
	if row-3 >= 0 && col-3 >= 0 &&
		grid[row-1][col-1] == "M" &&
		grid[row-2][col-2] == "A" &&
		grid[row-3][col-3] == "S" {
		count += 1
	}

	return count
}

func checkADiagonals(grid [][]string, row, col int) int {
	count := 0
	if row-1 >= 0 && col-1 >= 0 && row+1 < len(grid) && col+1 < len(grid[row]) {
		if grid[row-1][col-1] == "M" && grid[row-1][col+1] == "S" &&
			grid[row+1][col-1] == "M" && grid[row+1][col+1] == "S" { // example condition
			count += 1
		}

		if grid[row-1][col-1] == "S" && grid[row-1][col+1] == "M" &&
			grid[row+1][col-1] == "S" && grid[row+1][col+1] == "M" {
			count += 1
		}

		if grid[row-1][col-1] == "M" && grid[row-1][col+1] == "M" &&
			grid[row+1][col-1] == "S" && grid[row+1][col+1] == "S" {
			count += 1
		}

		if grid[row-1][col-1] == "S" && grid[row-1][col+1] == "S" &&
			grid[row+1][col-1] == "M" && grid[row+1][col+1] == "M" {
			count += 1
		}
	}
	return count
}
