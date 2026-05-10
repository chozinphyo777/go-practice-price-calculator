package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
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

func WriteJSON(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}
	file.Close()
	return nil
}
