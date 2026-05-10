package filemanager

import (
	"bufio"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Text is the line read from the file
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, err
	}
	file.Close()
	return lines, nil
}
