package repository

import (
	"fmt"
	"go-tts/internal/model"
)

// csvRegistryRepository handles the key-value mapping for Task Name -> Issue Key.
type csvRegistryRepository struct {
	filePath string
}

// NewRegistryRepository creates a new RegistryRepository.
func NewRegistryRepository(filePath string) RegistryRepository {
	return &csvRegistryRepository{filePath: filePath}
}

// LoadAll loads all registry entries from the CSV file.
func (r *csvRegistryRepository) LoadAll() ([]model.RegistryEntry, error) {
	records, err := ReadCSV(r.filePath)
	if err != nil {
		return nil, err
	}

	entries := make([]model.RegistryEntry, 0, len(records)-1)
	for i, record := range records {
		if i == 0 { // Skip header row
			continue
		}

		if len(record) != 4 {
			return nil, fmt.Errorf("invalid record length at line %d: expected 4, got %d", i+1, len(record))
		}
		entry := model.RegistryEntry{
			Task:          record[0],
			Issue:         record[1],
			Hyperlink:     record[2],
			MailIssueName: record[3],
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
