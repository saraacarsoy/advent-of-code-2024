package main

import (
	"log"
	"strings"
	utils "advent-of-code-2024/inputs"
	"strconv"
	"sort"
)

func main() {
	input, err := utils.ReadFile("day5")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	var rules, pages []string
	var isPagesSection bool

	for _, line := range input {
		if line == "" {
			isPagesSection = true
			continue
		}

		if isPagesSection {
			pages = append(pages, line)
		} else {
			rules = append(rules, line)
		}
	}

	rulesMap := formatRules(rules)

	checkPages(pages, rulesMap)
}

func checkPages(pages []string, rulesMap map[string][]string) {
	correctPages := 0
	for _, page := range pages {
		pageArr := strings.Split(page, ",")
		if isSortedAccordingToCustomOrder(pageArr, rulesMap) {
			// correctPages += getMiddleElement(pageArr)
		} else {
			customSort(pageArr, rulesMap)
			correctPages += getMiddleElement(pageArr)
		}
	}
	print(correctPages)
}

func isSortedAccordingToCustomOrder(input []string, rulesMap map[string][]string) bool {
	for i := 0; i < len(input)-1; i++ {
		a, b := input[i], input[i+1]

		if !isBefore(a, b, rulesMap) {
			return false
		}
	}

	return true
}

func isBefore(a, b string, rulesMap map[string][]string) bool {
	for _, rule := range rulesMap[a] {
		if rule == b {
			return true
		}
	}
	return false
}

func formatRules(rulesInput []string) map[string][]string {
	rulesMap := make(map[string][]string)

	for _, rule := range rulesInput {
		parts := strings.Split(rule, "|")

		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])

		if _, exists := rulesMap[parts[1]]; !exists {
			rulesMap[parts[1]] = []string{}
		}
	}

	return rulesMap
}

func getMiddleElement(arr []string) (int) {
	n := len(arr)

	middleElement := arr[n/2]

	middleElementInt, _ := strconv.Atoi(middleElement)

	return middleElementInt
}

func customSort(arr []string, rulesMap map[string][]string) {
	sort.Slice(arr, func(i, j int) bool {
		return isBefore(arr[i], arr[j], rulesMap)
	})
}
