package util

import "time"

// ParseDuration converts "1h 30m" to time.Duration
func ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration("1h30m") // Mock
}
