package repository

import (
	"encoding/csv"
	"os"
	"sync"
)

var fileMutex = &sync.Mutex{}

// ReadCSV reads a CSV file and returns its content as a slice of slices of strings.
func ReadCSV(path string) ([][]string, error) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteCSV writes a slice of slices of strings to a CSV file.
func WriteCSV(path string, data [][]string) error {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
