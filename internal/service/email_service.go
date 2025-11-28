package service

import (
	"bytes"
	"fmt"
	"go-tts/internal/model"
	"go-tts/internal/repository"
	"sort"
	"time"
)

type EmailService struct {
	ttsRepo      repository.TTSRepository
	registryRepo repository.RegistryRepository
}

func NewEmailService(ttsRepo repository.TTSRepository, registryRepo repository.RegistryRepository) *EmailService {
	return &EmailService{
		ttsRepo:      ttsRepo,
		registryRepo: registryRepo,
	}
}

func (s *EmailService) GenerateReport(date time.Time) (string, error) {
	// 1. Load all TTS rows.
	entries, err := s.ttsRepo.LoadAll()
	if err != nil {
		return "", err
	}

	// 2. Filter where Date == SelectedDate.
	var filteredEntries []model.TTSLogEntry
	for _, entry := range entries {
		if entry.Date.Year() == date.Year() && entry.Date.Month() == date.Month() && entry.Date.Day() == date.Day() {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	// 3. Group by Issue.
	groupedByIssue := make(map[string][]model.TTSLogEntry)
	for _, entry := range filteredEntries {
		groupedByIssue[entry.Issue] = append(groupedByIssue[entry.Issue], entry)
	}

	// 4. Fetch MailIssueName from Registry.
	registryEntries, err := s.registryRepo.LoadAll()
	if err != nil {
		return "", err
	}
	issueMailNameMap := make(map[string]string)
	for _, entry := range registryEntries {
		issueMailNameMap[entry.Issue] = entry.MailIssueName
	}

	// 5. Format the string.
	var report bytes.Buffer
	var issues []string
	for issue := range groupedByIssue {
		issues = append(issues, issue)
	}
	sort.Strings(issues)

	for _, issue := range issues {
		comments := groupedByIssue[issue]
		mailIssueName, ok := issueMailNameMap[issue]
		if !ok {
			mailIssueName = "General"
		}
		report.WriteString(fmt.Sprintf("(%s) %s:\n", issue, mailIssueName))
		for _, comment := range comments {
			report.WriteString(fmt.Sprintf("    - %s\n", comment.Comment))
		}
	}

	return report.String(), nil
}
