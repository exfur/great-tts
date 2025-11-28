package service

import (
	"fmt"
	"go-tts/internal/repository"
	"time"
)

type EmailService struct {
	ttsRepo repository.TTSRepository
}

func NewEmailService(r repository.TTSRepository) *EmailService {
	return &EmailService{ttsRepo: r}
}

func (s *EmailService) GenerateReport(date time.Time) string {
	// Logic to filter rows by date and format string
	return fmt.Sprintf("Report for %s\n...", date.Format("2006-01-02"))
}
