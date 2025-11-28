package service

import (
	"go-tts/internal/integration/jira"
	"go-tts/internal/repository"
)

type SyncService struct {
	ttsRepo    repository.TTSRepository
	jiraClient *jira.Client
}

func NewSyncService(repo repository.TTSRepository, client *jira.Client) *SyncService {
	return &SyncService{
		ttsRepo:    repo,
		jiraClient: client,
	}
}

func (s *SyncService) SyncApprovedWork() error {
	// 1. Load Repo
	// 2. Filter Ready & Not Logged
	// 3. Loop -> jiraClient.LogWork
	// 4. Update Repo IsLogged = true
	// 5. Save Repo
	return nil
}
