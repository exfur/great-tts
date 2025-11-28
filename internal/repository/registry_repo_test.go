package repository

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRegistryRepository_GetTaskIssueMap(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "registry-repo-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	filePath := filepath.Join(tmpdir, "registry.csv")
	csvData := `Task Name,Issue Key
Task A,PROJ-123
Task B,PROJ-456`

	if err := os.WriteFile(filePath, []byte(csvData), 0644); err != nil {
		t.Fatalf("Failed to write test CSV file: %v", err)
	}

	repo := NewRegistryRepository(filePath)
	taskMap, err := repo.GetTaskIssueMap()
	if err != nil {
		t.Fatalf("GetTaskIssueMap() failed: %v", err)
	}

	expectedMap := map[string]string{
		"Task A": "PROJ-123",
		"Task B": "PROJ-456",
	}

	if !reflect.DeepEqual(taskMap, expectedMap) {
		t.Errorf("GetTaskIssueMap() returned incorrect data.\nGot: %v\nWant: %v", taskMap, expectedMap)
	}
}
