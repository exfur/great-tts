package main

import (
	"log"
	"os"
	"path/filepath"

	"go-tts/config"
	"go-tts/internal/integration/jira"
	"go-tts/internal/repository"
	"go-tts/internal/service"
	"go-tts/internal/ui"

	"gioui.org/app"
)

func main() {
	// 1. Config
	cfg := config.Load()

	// 2. Data Layer
	ttsRepo := repository.NewTTSRepository(filepath.Join(cfg.DataDir, "tts.csv"))
	// Make sure this is defined
	registryRepo := repository.NewRegistryRepository(filepath.Join(cfg.DataDir, "registry.csv"))

	// 3. Integration Layer
	jiraClient := jira.NewClient(cfg.JiraBaseURL, cfg.JiraUser, cfg.JiraToken)

	// 4. Service Layer
	syncService := service.NewSyncService(ttsRepo, jiraClient)

	// Fix: Pass registryRepo here
	emailService := service.NewEmailService(ttsRepo, registryRepo)

	// 5. UI Layer
	// Fix: Pass registryRepo as the 3rd argument
	myApp := ui.NewUI(syncService, ttsRepo, registryRepo, emailService)

	// 6. Run
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Go TTS Manager"))

		if err := myApp.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
