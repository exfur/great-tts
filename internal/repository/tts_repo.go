package repository

import (
	"fmt"
	"go-tts/internal/model"
	"strconv"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

// csvTTSRepository implements the TTSRepository interface using a CSV file as storage.
type csvTTSRepository struct {
	filePath string
}

// NewTTSRepository creates a new TTSRepository that uses a CSV file.
func NewTTSRepository(filePath string) TTSRepository {
	return &csvTTSRepository{filePath: filePath}
}

// LoadAll retrieves all TTS log entries from the CSV file.
func (r *csvTTSRepository) LoadAll() ([]model.TTSLogEntry, error) {
	records, err := ReadCSV(r.filePath)
	if err != nil {
		return nil, err
	}

	entries := make([]model.TTSLogEntry, 0, len(records)-1)
	for i, record := range records {
		if i == 0 { // Skip header row
			continue
		}

		entry, err := r.parseRecord(record)
		if err != nil {
			return nil, fmt.Errorf("error parsing record at line %d: %w", i+1, err)
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// SaveAll saves all TTS log entries to the CSV file.
func (r *csvTTSRepository) SaveAll(entries []model.TTSLogEntry) error {
	records := make([][]string, 0, len(entries)+1)
	records = append(records, []string{"Date", "Task", "Comment", "From", "To", "Spent", "IsReady", "IsLogged", "Issue"})

	for _, entry := range entries {
		records = append(records, r.formatRecord(entry))
	}

	return WriteCSV(r.filePath, records)
}

func (r *csvTTSRepository) parseRecord(record []string) (model.TTSLogEntry, error) {
	var entry model.TTSLogEntry

	// Date,Task,Comment,From,To,Spent,IsReady,IsLogged,Issue
	if len(record) != 9 {
		return entry, fmt.Errorf("invalid record length: expected 9, got %d", len(record))
	}

	date, err := time.Parse(dateLayout, record[0])
	if err != nil {
		return entry, fmt.Errorf("failed to parse date: %w", err)
	}
	entry.Date = date

	entry.Task = record[1]
	entry.Comment = record[2]
	entry.From = record[3]
	entry.To = record[4]

	spent, err := time.ParseDuration(record[5])
	if err != nil {
		return entry, fmt.Errorf("failed to parse spent duration: %w", err)
	}
	entry.Spent = spent

	isReady, err := strconv.ParseBool(record[6])
	if err != nil {
		isReady = record[6] == "1"
	}
	entry.IsReady = isReady

	isLogged, err := strconv.ParseBool(record[7])
	if err != nil {
		isLogged = record[7] == "1"
	}
	entry.IsLogged = isLogged

	entry.Issue = record[8]

	return entry, nil
}

func (r *csvTTSRepository) formatRecord(entry model.TTSLogEntry) []string {
	isReady := "0"
	if entry.IsReady {
		isReady = "1"
	}

	isLogged := "0"
	if entry.IsLogged {
		isLogged = "1"
	}

	return []string{
		entry.Date.Format(dateLayout),
		entry.Task,
		entry.Comment,
		entry.From,
		entry.To,
		entry.Spent.String(),
		isReady,
		isLogged,
		entry.Issue,
	}
}
