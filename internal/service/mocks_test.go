package service

import (
	"go-tts/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockTTSRepository is a mock implementation of the TTSRepository interface.
type MockTTSRepository struct {
	mock.Mock
}

func (m *MockTTSRepository) LoadAll() ([]model.TTSLogEntry, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TTSLogEntry), args.Error(1)
}

func (m *MockTTSRepository) SaveAll(entries []model.TTSLogEntry) error {
	args := m.Called(entries)
	return args.Error(0)
}

func (m *MockTTSRepository) Save(entry model.TTSLogEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

// MockRegistryRepository is a mock implementation of the RegistryRepository interface.
type MockRegistryRepository struct {
	mock.Mock
}

func (m *MockRegistryRepository) LoadAll() ([]model.RegistryEntry, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.RegistryEntry), args.Error(1)
}

func (m *MockRegistryRepository) SaveAll(entries []model.RegistryEntry) error {
	args := m.Called(entries)
	return args.Error(0)
}

func (m *MockRegistryRepository) Save(entry model.RegistryEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}
