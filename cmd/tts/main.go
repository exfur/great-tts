package main

import (
	"log"
	"os"

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
	ttsRepo := repository.NewTTSRepo(cfg.DataDir)

	// 3. Integration Layer
	jiraClient := jira.NewClient(cfg.JiraBaseURL, cfg.JiraUser, cfg.JiraToken)

	// 4. Service Layer
	syncService := service.NewSyncService(ttsRepo, jiraClient)

	// 5. UI Layer
	myApp := ui.NewUI(syncService)

	// 6. Run
	go func() {
		// UPDATED: Create window instance directly
		w := new(app.Window)

		// Set window title
		w.Option(app.Title("Go TTS Manager"))

		if err := myApp.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
