package service

import (
	"fmt"
	"go-tts/internal/repository"
	"time"
)

type TTSService struct {
	ttsRepo      repository.TTSRepository
	registryRepo repository.RegistryRepository
}

func NewTTSService(ttsRepo repository.TTSRepository, registryRepo repository.RegistryRepository) *TTSService {
	return &TTSService{
		ttsRepo:      ttsRepo,
		registryRepo: registryRepo,
	}
}

func (s *TTSService) CalculateDuration(from, to string) (string, error) {
	fromTime, err := time.Parse("15:04", from)
	if err != nil {
		return "", fmt.Errorf("invalid 'from' time format: %w", err)
	}

	toTime, err := time.Parse("15:04", to)
	if err != nil {
		return "", fmt.Errorf("invalid 'to' time format: %w", err)
	}

	duration := toTime.Sub(fromTime)

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%dh %dm", hours, minutes), nil
}

func (s *TTSService) GetIssueForTask(taskName string) (string, error) {
	registryEntries, err := s.registryRepo.LoadAll()
	if err != nil {
		return "", err
	}

	for _, entry := range registryEntries {
		if entry.Task == taskName {
			return entry.Issue, nil
		}
	}

	return "", nil
}
