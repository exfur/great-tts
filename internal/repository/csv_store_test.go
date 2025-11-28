package repository

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadWriteCSV(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "csv-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	filePath := filepath.Join(tmpdir, "test.csv")
	testData := [][]string{
		{"header1", "header2"},
		{"value1", "value2"},
		{"value3", "value4"},
	}

	err = WriteCSV(filePath, testData)
	if err != nil {
		t.Fatalf("WriteCSV failed: %v", err)
	}

	readData, err := ReadCSV(filePath)
	if err != nil {
		t.Fatalf("ReadCSV failed: %v", err)
	}

	if !reflect.DeepEqual(testData, readData) {
		t.Errorf("Read data does not match written data.\nGot: %v\nWant: %v", readData, testData)
	}
}
