package repository

import "go-tts/internal/model"

type TTSRepository interface {
	LoadAll() ([]model.TTSLogEntry, error)
	SaveAll([]model.TTSLogEntry) error
	Save(model.TTSLogEntry) error
}

type RegistryRepository interface {
	LoadAll() ([]model.RegistryEntry, error)
	SaveAll([]model.RegistryEntry) error
	Save(model.RegistryEntry) error
}
