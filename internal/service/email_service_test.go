package service

import (
	"errors"
	"go-tts/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmailService_GenerateReport(t *testing.T) {
	mockTTSRepo := new(MockTTSRepository)
	mockRegistryRepo := new(MockRegistryRepository)
	service := NewEmailService(mockTTSRepo, mockRegistryRepo)

	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	entries := []model.TTSLogEntry{
		{Date: date, Task: "Task 1", Comment: "Comment 1", Issue: "Issue 1"},
		{Date: date, Task: "Task 2", Comment: "Comment 2", Issue: "Issue 1"},
		{Date: date, Task: "Task 3", Comment: "Comment 3", Issue: "Issue 2"},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Task: "Task 4", Comment: "Comment 4", Issue: "Issue 3"},
	}

	registryEntries := []model.RegistryEntry{
		{Task: "Task 1", Issue: "Issue 1", MailIssueName: "Mail Issue 1"},
		{Task: "Task 3", Issue: "Issue 2", MailIssueName: "Mail Issue 2"},
	}

	mockTTSRepo.On("LoadAll").Return(entries, nil)
	mockRegistryRepo.On("LoadAll").Return(registryEntries, nil)

	report, err := service.GenerateReport(date)
	assert.NoError(t, err)

	expectedReport := "(Issue 1) Mail Issue 1:\n    - Comment 1\n    - Comment 2\n(Issue 2) Mail Issue 2:\n    - Comment 3\n"
	assert.Equal(t, expectedReport, report)

	mockTTSRepo.AssertExpectations(t)
	mockRegistryRepo.AssertExpectations(t)
}

func TestEmailService_GenerateReport_ErrorLoadingEntries(t *testing.T) {
	mockTTSRepo := new(MockTTSRepository)
	mockRegistryRepo := new(MockRegistryRepository)
	service := NewEmailService(mockTTSRepo, mockRegistryRepo)

	mockTTSRepo.On("LoadAll").Return([]model.TTSLogEntry(nil), errors.New("error loading entries"))

	_, err := service.GenerateReport(time.Now())
	assert.Error(t, err)

	mockTTSRepo.AssertExpectations(t)
}
