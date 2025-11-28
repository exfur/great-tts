package repository

import "go-tts/internal/model"

type TTSRepository interface {
	LoadAll() ([]model.TTSLogEntry, error)
	SaveAll([]model.TTSLogEntry) error
}

type RegistryRepository interface {
	LoadAll() ([]model.RegistryEntry, error)
}
