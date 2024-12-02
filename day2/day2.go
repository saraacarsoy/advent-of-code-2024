package main

import (
	utils "AoC2024/inputs"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("day2")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	var count int

	for _, line := range input {
		parts := strings.Fields(line)

		var intParts []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			intParts = append(intParts, num)
		}

		if checkOrder(intParts) {
			count += 1
		}
	}
	print(count)
}

func checkOrder(baseArray []int) bool {
	return isSortedAscending(baseArray) || isSortedDescending(baseArray)
}

func isSortedAscending(slice []int) bool {
	var controlArr []bool

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			return false
		}

		var correctDecrease bool = slice[i]-slice[i+1] == 3 || slice[i]-slice[i+1] == 2 || slice[i]-slice[i+1] == 1
		controlArr = append(controlArr, correctDecrease)
	}
	return allTrue(controlArr)
}

func isSortedDescending(slice []int) bool {
	var controlArr []bool
	fmt.Println(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}

		var correctIncrease bool = slice[i]-slice[i+1] == -3 || slice[i]-slice[i+1] == -2 || slice[i]-slice[i+1] == -1
		controlArr = append(controlArr, correctIncrease)
	}

	fmt.Println(controlArr)
	return allTrue(controlArr)
}

func allTrue(slice []bool) bool {
	for _, val := range slice {
		if !val {
			return false
		}
	}
	return true
}
