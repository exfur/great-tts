package repository

import (
	"go-tts/internal/model"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestTTSRepository_LoadAll(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "tts-repo-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	filePath := filepath.Join(tmpdir, "tts.csv")
	csvData := `Date,Task,Comment,From,To,Spent,IsReady,IsLogged,Issue
2023-10-27,Task 1,Comment 1,10:00,11:00,1h0m0s,1,0,ISSUE-1
2023-10-28,Task 2,Comment 2,14:00,15:30,1h30m0s,0,1,ISSUE-2`

	if err := os.WriteFile(filePath, []byte(csvData), 0644); err != nil {
		t.Fatalf("Failed to write test CSV file: %v", err)
	}

	repo := NewTTSRepository(filePath)
	entries, err := repo.LoadAll()
	if err != nil {
		t.Fatalf("LoadAll() failed: %v", err)
	}

	expectedDate1, _ := time.Parse("2006-01-02", "2023-10-27")
	expectedDate2, _ := time.Parse("2006-01-02", "2023-10-28")
	expectedDuration1, _ := time.ParseDuration("1h")
	expectedDuration2, _ := time.ParseDuration("1h30m")

	expectedEntries := []model.TTSLogEntry{
		{Date: expectedDate1, Task: "Task 1", Comment: "Comment 1", From: "10:00", To: "11:00", Spent: expectedDuration1, IsReady: true, IsLogged: false, Issue: "ISSUE-1"},
		{Date: expectedDate2, Task: "Task 2", Comment: "Comment 2", From: "14:00", To: "15:30", Spent: expectedDuration2, IsReady: false, IsLogged: true, Issue: "ISSUE-2"},
	}

	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("LoadAll() returned incorrect data.\nGot: %+v\nWant: %+v", entries, expectedEntries)
	}
}
