package model

import "time"

type TTSLogEntry struct {
	Date      time.Time
	Task      string
	Comment   string
	From      string
	To        string
	Spent     string // Raw string from CSV
	Issue     string
	IsReady   bool
	IsLogged  bool
	LogErrors string
}
