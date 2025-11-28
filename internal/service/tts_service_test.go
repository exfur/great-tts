package service

import (
	"go-tts/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestTTSService_CalculateDuration(t *testing.T) {
	service := NewTTSService(nil, nil)

	testCases := []struct {
		name     string
		from     string
		to       string
		expected string
		hasError bool
	}{
		{"Valid duration", "10:00", "11:30", "1h 30m", false},
		{"Invalid from time", "invalid", "11:30", "", true},
		{"Invalid to time", "10:00", "invalid", "", true},
		{"Zero duration", "10:00", "10:00", "0h 0m", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			duration, err := service.CalculateDuration(tc.from, tc.to)

			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, duration)
			}
		})
	}
}

func TestTTSService_GetIssueForTask(t *testing.T) {
	mockRegistryRepo := new(MockRegistryRepository)
	service := NewTTSService(nil, mockRegistryRepo)

	registryEntries := []model.RegistryEntry{
		{Task: "Task 1", Issue: "Issue 1"},
		{Task: "Task 2", Issue: "Issue 2"},
	}

	mockRegistryRepo.On("LoadAll").Return(registryEntries, nil)

	issue, err := service.GetIssueForTask("Task 1")
	assert.NoError(t, err)
	assert.Equal(t, "Issue 1", issue)

	issue, err = service.GetIssueForTask("Non-existent task")
	assert.NoError(t, err)
	assert.Equal(t, "", issue)

	mockRegistryRepo.AssertExpectations(t)
}
