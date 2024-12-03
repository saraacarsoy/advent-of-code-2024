package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(day string) ([]string, error) {
	baseDir := `C:\Users\nilac\GoProjects\advent-of-code-2024\inputs`

	filePath := filepath.Join(baseDir, fmt.Sprintf("%s.txt", day))

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
