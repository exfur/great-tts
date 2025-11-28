package repository

import (
	"fmt"
)

// RegistryRepository handles the key-value mapping for Task Name -> Issue Key.
type RegistryRepository struct {
	filePath string
}

// NewRegistryRepository creates a new RegistryRepository.
func NewRegistryRepository(filePath string) *RegistryRepository {
	return &RegistryRepository{filePath: filePath}
}

// GetTaskIssueMap loads the Task Name -> Issue Key mapping from the CSV file.
func (r *RegistryRepository) GetTaskIssueMap() (map[string]string, error) {
	records, err := ReadCSV(r.filePath)
	if err != nil {
		return nil, err
	}

	taskIssueMap := make(map[string]string)
	for i, record := range records {
		if i == 0 { // Skip header row
			continue
		}

		if len(record) != 2 {
			return nil, fmt.Errorf("invalid record length at line %d: expected 2, got %d", i+1, len(record))
		}
		taskName := record[0]
		issueKey := record[1]
		taskIssueMap[taskName] = issueKey
	}

	return taskIssueMap, nil
}
