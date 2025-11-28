package config

type Config struct {
	JiraBaseURL string
	JiraUser    string
	JiraToken   string
	DataDir     string
}

func Load() *Config {
	return &Config{
		JiraBaseURL: "https://jira.glowbyteconsulting.com",
		DataDir:     "./data",
	}
}
