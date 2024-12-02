package main

import (
	utils "AoC2024/inputs"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("day1")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	var firstList []int
	var secondList []int

	for _, line := range input {
		parts := strings.Fields(line)

		firstNum, secondNum := parts[0], parts[1]

		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)

		firstList = append(firstList, firstNumInt)
		secondList = append(secondList, secondNumInt)
	}
	// part1(firstList, secondList)

	var total int
	for _, i := range firstList {
		amount := countOccurrences(secondList, i)
		total += i * amount
	}
	fmt.Println(total)
}

func part1(firstList []int, secondList []int) {
	sort.Ints(firstList)
	sort.Ints(secondList)

	var total int
	for i := range firstList {
		var diff int = (firstList[i]) - (secondList[i])
		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	fmt.Println(total)
}

func countOccurrences(slice []int, num int) int {
	count := 0
	for _, v := range slice {
		if v == num {
			count++
		}
	}
	return count
}
