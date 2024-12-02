package main

import (
	utils "AoC2024/inputs"
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
		} else { // part 2
			for i := 0; i < len(intParts); i++ {
				newArr := make([]int, len(intParts))
				copy(newArr, intParts)

				newArr = append(newArr[:i], newArr[i+1:]...)
				if checkOrder(newArr) {
					count += 1
					break
				}
			}
		}
	}

	print(count)
}

func checkOrder(baseArray []int) bool {
	return isSortedAscending(baseArray) || isSortedDescending(baseArray)
}

func isSortedDescending(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			return false
		}

		if !(slice[i]-slice[i+1] == 3 || slice[i]-slice[i+1] == 2 || slice[i]-slice[i+1] == 1) {
			return false
		}
	}
	return true
}

func isSortedAscending(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}

		if !(slice[i]-slice[i+1] == -3 || slice[i]-slice[i+1] == -2 || slice[i]-slice[i+1] == -1) {
			return false
		}
	}
	return true
}
