package model

import "time"

type TTSLogEntry struct {
	Date      time.Time
	Task      string
	Comment   string
	From      string
	To        string
	Spent     time.Duration
	Issue     string
	IsReady   bool
	IsLogged  bool
	LogErrors string
}
