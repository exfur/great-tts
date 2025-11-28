package repository

import "go-tts/internal/model"

type TTSRepository interface {
	LoadAll() ([]model.TTSLogEntry, error)
	SaveAll([]model.TTSLogEntry) error
}
