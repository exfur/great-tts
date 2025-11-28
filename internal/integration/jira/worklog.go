package jira

import "time"

type WorklogPayload struct {
	Comment          string `json:"comment"`
	Started          string `json:"started"`
	TimeSpentSeconds int    `json:"timeSpentSeconds"`
}

func (c *Client) LogWork(issueKey string, date time.Time, durationSec int, comment string) error {
	// HTTP POST logic to Jira REST API
	return nil
}
