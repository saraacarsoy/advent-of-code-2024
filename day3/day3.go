package main

import (
	utils "advent-of-code-2024/inputs"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input, err := utils.ReadFile("day3")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// part1(input)
	part2(input)
}

func part1(input []string) {
	pattern := `mul\(([0-9]+),([0-9]+)\)`
	re := regexp.MustCompile(pattern)

	count := 0
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			count += left * right
		}
	}
	print(count)
}

func part2(input []string) {
	shouldMultiply := true
	pattern := `mul\(([0-9]+),([0-9]+)\)|don't\(\)|do\(\)`
	re := regexp.MustCompile(pattern)

	count := 0
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) >= 3 && match[1] != "" && match[2] != "" && shouldMultiply {
				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])

				count += left * right
			} else {
				if match[0] == "don't()" {
					shouldMultiply = false
				} else if match[0] == "do()" {
					shouldMultiply = true
				}
			}
		}
	}

	print(count)
}
